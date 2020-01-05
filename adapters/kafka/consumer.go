package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	kafka "github.com/bsm/sarama-cluster"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

// ConsumerCallback define callback
type ConsumerCallback func([]byte)

// ComsumerHandler read msg
type ComsumerHandler interface {
	Read(cb ConsumerCallback)
}

type comsumerHandler struct {
	consumer *kafka.Consumer
	topics   []string
	group    string
}

// CreateConsumers init consume
func CreateConsumers(cfg *Kafka) ComsumerHandler {
	config := kafka.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	if cfg.Newest {
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
	} else {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	if cfg.PemFiles != nil {
		tls, err := NewTLSConfig(*cfg.PemFiles)
		if err != nil {
			logrus.Errorf("err create tls config: %v", err)
			return &comsumerHandler{}
		}
		tls.InsecureSkipVerify = cfg.InsecureSkipVerify
		config.Net.TLS.Enable = true
		config.Net.TLS.Config = tls
	}

	c, err := kafka.NewConsumer(cfg.Addrs, cfg.Group, cfg.Topics, config)

	if err != nil {
		logrus.Errorf("err new consume: %v", err)
		return &comsumerHandler{}
	}

	go func() {
		for err := range c.Errors() {
			logrus.Errorf("err consume: %v", err)
		}
	}()

	go func() {
		for ntf := range c.Notifications() {
			logrus.Error("consume notifications: %v", ntf)
		}
	}()
	return &comsumerHandler{consumer: c, topics: cfg.Topics, group: cfg.Group}
}

// Read msg when consume
func (c *comsumerHandler) Read(cb ConsumerCallback) {
	if c.consumer == nil {
		logrus.Error(context.Background(), "consumer is nil")
		return
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	consumer := c.consumer
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				cb(msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case <-signals:
			return
		}
	}
}
