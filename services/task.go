package services

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"todolist-facebook-chatbot/adapters/kafka"
	"todolist-facebook-chatbot/conf"
	"todolist-facebook-chatbot/dtos"
	"todolist-facebook-chatbot/models"
)

// Task services
type ITaskService interface {
	Create(request *dtos.CreateTaskRequest) (*dtos.CreateTaskResponse, error)
}

type taskService struct {
	kafkaAdapt kafka.ProducerV2
	config     conf.AppConfig
}

func NewTaskService(
	kafkaAdapter kafka.ProducerV2,
	config conf.AppConfig,
) ITaskService {
	return &taskService{
		kafkaAdapt: kafkaAdapter,
		config:     config,
	}
}

func (t *taskService) Create(request *dtos.CreateTaskRequest) (*dtos.CreateTaskResponse, error) {
	task := &models.Task{
		Title:       request.Title,
		Description: request.Description,
		StartAt:     request.StartAt,
		EndAt:       request.EndAt,
	}
	//err := t.taskRepo.Create(task)
	//if err != nil {
	//	logrus.Fatalln("Error when creating task", err)
	//	return nil, err
	//}
	if t.kafkaAdapt != nil {
		err := t.kafkaAdapt.WriteByTopic(t.convertToKafkaMsg(task), t.config.Name)
		if err != nil {
			logrus.Error("Got error while push kafka msg: %v", err)
		}
	}
	return &dtos.CreateTaskResponse{
		Meta: dtos.NewMeta(http.StatusOK),
	}, nil
}

func (t *taskService) convertToKafkaMsg(task *models.Task) *dtos.TaskKafkaMsg {
	return &dtos.TaskKafkaMsg{
		ID:          "",
		RequestID:   "",
		RefEventID:  "",
		Event:       "",
		ServiceCode: "",
		TimeStamp:   0,
		UserID:      "",
		PayloadID:   "",
		Payload: dtos.TaskKafkaPayload{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			StartAt:     task.StartAt,
			EndAt:       task.EndAt,
		},
	}
}