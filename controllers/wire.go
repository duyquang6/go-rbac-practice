// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"todolist-facebook-chatbot/conf"
	"todolist-facebook-chatbot/repositories"
	"todolist-facebook-chatbot/services"
)

func injectTaskService() (*services.TaskService, error) {
	wire.Build(conf.NewAppConfig, repositories.NewResource, repositories.NewTaskRepository, services.NewTaskService)
	return &services.TaskService{}, nil
}