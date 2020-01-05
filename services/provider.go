package services

import (
	"go.uber.org/dig"
	"strings"
	_kafka "todolist-facebook-chatbot/adapters/kafka"
	"todolist-facebook-chatbot/conf"
)

// serviceContainer is a global ServiceProvider.
var serviceContainer *dig.Container

func InitializeServices() {
	serviceContainer = dig.New()
	_ = serviceContainer.Provide(conf.NewAppConfig)

	_ = serviceContainer.Provide(func(config conf.AppConfig) _kafka.ProducerV2 {
		var (
			addrConfig = config.Kafka.Address
		)

		kafkaAdders := strings.Split(addrConfig, ",")
		cfg := &_kafka.Kafka{
			Addrs: kafkaAdders,
		}
		return _kafka.CreateWriters(cfg)
	})

	//_ = serviceContainer.Provide(repositories.NewResource)

	//_ = serviceContainer.Provide(repositories.NewTaskRepository)

	_ = serviceContainer.Provide(NewTaskService)
}

func GetServiceContainer() *dig.Container {
	return serviceContainer
}
