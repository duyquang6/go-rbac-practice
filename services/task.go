package services

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"net/http"
	"todolist-facebook-chatbot/conf"
	"todolist-facebook-chatbot/dtos"
	"todolist-facebook-chatbot/models"
	"todolist-facebook-chatbot/repositories"
)

// Task services
type ITaskService interface {
	Create(request *dtos.CreateTaskRequest) (*dtos.CreateTaskResponse, error)
}

var provideSet = wire.NewSet(conf.NewAppConfig, repositories.NewResource, repositories.NewTaskRepository, NewTaskService)

type TaskService struct {
	taskRepo repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) (*TaskService, error) {
	return &TaskService{
		taskRepo: taskRepo,
	}, nil
}

func (t TaskService) Create(request *dtos.CreateTaskRequest) (*dtos.CreateTaskResponse, error) {
	logrus.Println("Creating Task")
	err := t.taskRepo.Create(&models.Task{
		Title:       request.Title,
		Description: request.Description,
		StartAt:     request.StartAt,
		EndAt:       request.EndAt,
	})
	if err != nil {
		logrus.Fatalln("Error when creating task", err)
	}
	return &dtos.CreateTaskResponse{
		Meta: dtos.NewMeta(http.StatusOK),
	}, nil
}
