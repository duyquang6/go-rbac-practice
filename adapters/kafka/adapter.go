package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

// Const for Kafka adapter
const (
	TimeOut time.Duration = 15
)

// Adapter the adapter to comunicate via kafka
type Adapter interface {
	Close()
	GetProducer() *Producer
	CreateTopic(numPartition int32) error
	Publish(content string) error
	PublishMessage(content interface{}, eventType int) error
}

type adapter struct {
	producer *Producer
	host     string
	port     int32
	topic    string
}

// Message represent the kafka message
type Message struct {
	Data      interface{} `json:"data"`
	EventType int         `json:"event_type"`
}

// NewKafkaAdapter create the new kafka client
func NewKafkaAdapter(host string, port int32, topic string) (Adapter, error) {
	kafkaSeedBroker := fmt.Sprintf("%v:%v", host, port)
	kafkaProducer, err := NewKafkaProducer(&ProducerConfig{
		SeedBrokers:      []string{kafkaSeedBroker},
		NumFlushMessages: 1,
	})
	if err != nil {
		return nil, err
	}

	return &adapter{
		producer: kafkaProducer,
		topic:    topic,
		host:     host,
		port:     port,
	}, nil
}

func (a *adapter) CreateTopic(numPartition int32) error {
	broker := sarama.NewBroker(fmt.Sprintf("%v:%v", a.host, a.port))
	config := sarama.NewConfig()
	config.Version = sarama.V1_1_0_0
	err := broker.Open(config)
	if err != nil {
		return err
	}

	// Setup the Topic details in CreateTopicRequest struct
	topicDetail := &sarama.TopicDetail{}
	topicDetail.NumPartitions = numPartition
	topicDetail.ReplicationFactor = int16(1)
	topicDetail.ConfigEntries = make(map[string]*string)

	topicDetails := make(map[string]*sarama.TopicDetail)
	topicDetails[a.topic] = topicDetail

	request := sarama.CreateTopicsRequest{
		Timeout:      time.Second * TimeOut,
		TopicDetails: topicDetails,
	}

	// Send request to Broker
	response, err := broker.CreateTopics(&request)
	if err != nil {
		return err
	}

	t := response.TopicErrors
	for key, val := range t {
		if strings.HasSuffix(val.Err.Error(), "exists.") {
			logrus.Info("%v (topic: '%v')", val.Err.Error(), key)
		}

		if strings.HasSuffix(val.Err.Error(), "why are you printing me?") {
			logrus.Info("created kafka topic: '%v'", key)
		}
	}

	// close connection to broker
	return broker.Close()
}

func (a *adapter) Publish(content string) error {
	a.producer.SendMessage(&ProducerMessage{
		Topic: a.topic,
		Value: []byte(content),
	})
	return nil
}

func (a *adapter) PublishMessage(content interface{}, eventType int) error {
	jsonMsg := &Message{
		Data:      content,
		EventType: eventType,
	}

	msgBytes, err := json.Marshal(jsonMsg)
	if err != nil {
		return err
	}
	a.producer.SendMessage(&ProducerMessage{
		Topic: a.topic,
		Key:   nil,
		Value: msgBytes,
	})

	return nil
}

func (a *adapter) Close() {
	a.producer.Close()
}

func (a *adapter) GetProducer() *Producer {
	return a.producer
}
