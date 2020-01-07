package services

import (
	"go.uber.org/dig"
	"strings"
	_kafka "todolist-facebook-chatbot/adapters/kafka"
	"todolist-facebook-chatbot/adapters/slack"
	"todolist-facebook-chatbot/conf"
)

// serviceContainer is a global ServiceProvider.
var serviceContainer *dig.Container

func init() {
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

	_ = serviceContainer.Provide(func(config conf.AppConfig) slack.Adapter {
		var (
			webhookURL = config.Slack.WebHookURL
		)

		return slack.NewSlackAPI(webhookURL)
	})

	//_ = serviceContainer.Provide(repositories.NewResource)

	//_ = serviceContainer.Provide(repositories.NewTaskRepository)

	_ = serviceContainer.Provide(NewTaskService)
}

func GetServiceContainer() *dig.Container {
	return serviceContainer
}
