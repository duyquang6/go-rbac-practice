package conf

import (
	"fmt"
)

type Postgres struct {
	Username string `default:"postgres" envconfig:"USERNAME"`
	Password string `default:"1234" envconfig:"PASSWORD"`
	Host     string `default:"localhost" envconfig:"HOST"`
	Port     int    `default:"5432" envconfig:"PORT"`
	Database string `default:"todolist_facebook_chatbot" envconfig:"DATABASE"`
}

// ConnectionString returns connection string of Postgres database.
func (c *Postgres) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		c.Host, c.Port, c.Username, c.Database, c.Password)
}
