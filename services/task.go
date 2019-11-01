package services

import (
	"context"
	"log"
	"net/http"
	"todolist-facebook-chatbot/dtos"
)

// Task services
type TaskService interface {
	Create(ctx context.Context, request *dtos.CreateTaskRequest) (*dtos.CreateTaskResponse, error)
}

type taskService struct {
}

// NewNtfService returns a new instance of Merchant QR Service.
func NewTaskService() TaskService {
	return &taskService{}
}

func (t taskService) Create(ctx context.Context, request *dtos.CreateTaskRequest) (*dtos.CreateTaskResponse, error) {
	log.Println("Creating Task")

	return &dtos.CreateTaskResponse{
		Meta: dtos.NewMeta(http.StatusOK),
	}, nil
}
