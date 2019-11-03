package conf

import (
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Host    string   `default:"0.0.0.0" envconfig:"HOST"`
	Port    int      `default:"8080" envconfig:"PORT"`
	RunMode string   `default:"debug" envconfig:"RUN_MODE"`
	DB      Postgres `envconfig:"DB"`
}

// NewNtfServiceConfig create new notification service.
func NewAppConfig() (*AppConfig, error) {
	var c AppConfig

	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
