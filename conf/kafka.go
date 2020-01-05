package conf

type Kafka struct {
	Address  string `default:"0.0.0.0:9092" envconfig:"ADDRESS"`
	UseTLS   bool   `default:"false" envconfig:"TLS_ENABLED"`

	// Topic
	Topic
}

type Topic struct {
	CreateTask string `default:"task_created" envconfig:"TASK_CREATE"`
}
