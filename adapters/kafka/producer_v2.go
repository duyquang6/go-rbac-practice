package kafka

import (
	"context"
	"encoding/json"
	kafka "github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"time"
)

// ProducerV2 interface
type ProducerV2 interface {
	WriteRaw([]byte)
	Write(kafka.Encoder)
	WriteRawByTopic([]byte, string)
	Close()
	WriteByTopic(interface{}, string) error
}
type producerV2 struct {
	topics   []string
	producer kafka.AsyncProducer
}

// CreateWriters producer
func CreateWriters(cfgKafka *Kafka) ProducerV2 {
	cfg := kafka.NewConfig()
	cfg.Producer.RequiredAcks = kafka.WaitForLocal
	cfg.Producer.Flush.Frequency = 50 * time.Millisecond
	if cfgKafka.MaxMessageBytes > 0 {
		cfg.Producer.MaxMessageBytes = cfgKafka.MaxMessageBytes
	}
	if cfgKafka.Compress {
		cfg.Producer.Compression = kafka.CompressionGZIP
	}
	if cfgKafka.PemFiles != nil {
		tls, err := NewTLSConfig(*cfgKafka.PemFiles)
		if err != nil {
			logrus.Error(context.Background(), "err create tls config: %v", err)
			return nil
		}
		cfg.Net.TLS.Enable = true
		cfg.Net.TLS.Config = tls
	}

	producer, err := kafka.NewAsyncProducer(cfgKafka.Addrs, cfg)
	if err != nil {
		logrus.Error(context.Background(), "Failed to write entry: %v,%v,%v", cfgKafka.Addrs, cfgKafka.Topics, err)
		return nil
	}
	go func() {

		for err := range producer.Errors() {
			logrus.Error(context.Background(), "Failed to write entry: %v", err)
		}
	}()
	return &producerV2{
		topics:   cfgKafka.Topics,
		producer: producer,
	}
}

func (w *producerV2) Write(v kafka.Encoder) {
	if w == nil {
		logrus.Error(context.Background(), "Write is nil")
		return
	}
	for _, topic := range w.topics {
		w.producer.Input() <- &kafka.ProducerMessage{
			Topic: topic,
			Value: v,
		}
	}
}

// Close kafka
func (w *producerV2) Close() {
	w.producer.AsyncClose()
}

// WriteRaw kafka
func (w *producerV2) WriteRaw(v []byte) {
	if w == nil {
		logrus.Error(context.Background(), "Write is nil")
		return
	}
	for _, topic := range w.topics {
		w.producer.Input() <- &kafka.ProducerMessage{
			Topic: topic,
			Value: kafka.ByteEncoder(v),
		}
	}
}

// WriteRawByTopic kafka
func (w *producerV2) WriteRawByTopic(v []byte, topicName string) {
	if w == nil {
		logrus.Error(context.Background(), "Write is nil")
		return
	}
	w.producer.Input() <- &kafka.ProducerMessage{
		Topic: topicName,
		Value: kafka.ByteEncoder(v),
	}
}

// WriteByTopic kafka
func (w *producerV2) WriteByTopic(v interface{}, topicName string) error {
	data, err := json.Marshal(v)
	if err != nil {
		logrus.Error(context.Background(), "Marshal error: %v", err)
		return err
	}
	w.producer.Input() <- &kafka.ProducerMessage{
		Topic: topicName,
		Value: kafka.ByteEncoder(data),
	}
	return nil
}
