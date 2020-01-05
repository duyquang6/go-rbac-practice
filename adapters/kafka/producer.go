package kafka

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"reflect"
	"sync"
	"time"
)

// ProducerType type of producer
type ProducerType string

// ProducerType const
const (
	KafkaSyncProducerType  ProducerType = "sync"
	KafkaAsyncProducerType ProducerType = "async"
)

// Error variables define
var (
	ErrNoTopicDefined = errors.New("no topic defined")
)

// ProducerMessage message of producer
type ProducerMessage struct {
	Topic string
	Key   []byte
	Value []byte
}

// ProducerConfig the kafka adapter
type ProducerConfig struct {
	SeedBrokers      []string
	NumFlushMessages int
	TopicMap         map[string]string
}

// Producer struct represent producer of kafka
type Producer struct {
	producer sarama.AsyncProducer
	topicMap map[string]string
}

// NewKafkaProducer create the new producer
func NewKafkaProducer(cf *ProducerConfig) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Flush.Messages = cf.NumFlushMessages
	config.Producer.Flush.Frequency = 1 * time.Second
	config.Producer.Return.Successes = false
	config.Producer.Return.Errors = false

	asyncProducer, err := sarama.NewAsyncProducer(cf.SeedBrokers, config)
	if err != nil {
		return nil, err
	}

	kafkaProducer := &Producer{
		producer: asyncProducer,
		topicMap: cf.TopicMap,
	}

	return kafkaProducer, nil
}

// SendMessage send the message to kafka queue
func (p *Producer) SendMessage(m *ProducerMessage) {
	msg := &sarama.ProducerMessage{
		Topic: m.Topic,
		Key:   sarama.ByteEncoder(m.Key),
		Value: sarama.ByteEncoder(m.Value),
	}

	p.producer.Input() <- msg
}

// SendAbstractMessage send the astract message to kafka queue
func (p *Producer) SendAbstractMessage(msg interface{}) error {
	msgStructName := p.getTypeOfMessage(msg)
	topic := p.topicMap[msgStructName]
	if topic == "" {
		return ErrNoTopicDefined
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   nil,
		Value: sarama.ByteEncoder(msgBytes),
	}

	p.producer.Input() <- kafkaMsg
	return nil
}

func (p *Producer) getTypeOfMessage(msg interface{}) string {
	t := reflect.TypeOf(msg)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}

// Close the producer
func (p *Producer) Close() {
	var wg sync.WaitGroup
	p.producer.AsyncClose()

	wg.Add(2)
	go func() {
		for range p.producer.Successes() {
			fmt.Println("Unexpected message on Successes()")
		}
		wg.Done()
	}()
	go func() {
		for msg := range p.producer.Errors() {
			fmt.Println(msg.Err)
		}
		wg.Done()
	}()
	wg.Wait()
}
