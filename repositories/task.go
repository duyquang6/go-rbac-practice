package repositories

import (
	"todolist-facebook-chatbot/conf"
	"todolist-facebook-chatbot/models"
)

// CachedRepository provides access to a Redis caching database.
type TaskRepository interface {
	Create(*models.Task) error
	GetAll() ([]*models.Task, error)
}

type taskRepository struct {
	resource  Resource
	appConfig conf.AppConfig
}

func NewTaskRepository(config conf.AppConfig, resource Resource) TaskRepository {
	return &taskRepository{
		resource:  resource,
		appConfig: config,
	}
}

func (repo taskRepository) Create(task *models.Task) error {
	return repo.resource.DB.Create(task).Error
}

func (r taskRepository) GetAll() ([]*models.Task, error) {
	tasks := []*models.Task{}
	err := r.resource.DB.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
