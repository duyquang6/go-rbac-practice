package repositories

import (
	"log"
	"todolist-facebook-chatbot/conf"
	"todolist-facebook-chatbot/models"
)

// CachedRepository provides access to a Redis caching database.
type TaskRepository interface {
	Create(*models.Task) error
}

type taskRepository struct {
	resource  *Resource
	appConfig *conf.AppConfig
}

func NewTaskRepository(config *conf.AppConfig, resource *Resource) (TaskRepository, error) {
	return &taskRepository{
		resource:  resource,
		appConfig: config,
	}, nil
}

func (repo taskRepository) Create(task *models.Task) error {
	log.Println("THE HELL")
	return nil
}
