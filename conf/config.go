package conf

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	Host    string   `default:"0.0.0.0" envconfig:"HOST"`
	Port    int      `default:"8080" envconfig:"PORT"`
	RunMode string   `default:"debug" envconfig:"RUN_MODE"`
	Name    string   `default:"todolist" envconfig:"APP_NAME"`
	DB      Postgres `envconfig:"DB"`
	Kafka   Kafka    `envconfig:"KAFKA"`
}

// NewNtfServiceConfig create new notification service.
func NewAppConfig() AppConfig {
	var c AppConfig
	if err := envconfig.Process("", &c); err != nil {
		logrus.Error("Cannot load config %v", err)
	}
	return c
}
