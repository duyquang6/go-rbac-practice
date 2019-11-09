package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"todolist-facebook-chatbot/conf"
	"todolist-facebook-chatbot/models"
)

type Resource struct {
	DB *gorm.DB
}

func autoMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Task{})
}

func NewResource(cfg *conf.AppConfig) (*Resource, error) {
	log.Println("Initializing DB . . .")
	db, err := gorm.Open("postgres", cfg.DB.GetConnectionString())
	if err != nil {
		log.Fatalln("Got error when init DB: ", err)
	}

	autoMigration(db)

	return &Resource{DB: db}, err
}
