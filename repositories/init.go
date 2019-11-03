package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"todolist-facebook-chatbot/conf"
)

type Resource struct {
	DB *gorm.DB
}

func NewResource(cfg *conf.AppConfig) (*Resource, error) {
	log.Println("Initializing DB . . .")
	db, err := gorm.Open("postgres", cfg.DB.GetConnectionString())
	defer db.Close()
	if err != nil {
		log.Fatalln("Got error when init DB: ", err)
	}
	
	return &Resource{DB: db}, err
}