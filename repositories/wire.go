// +build wireinject

package repositories

import (
	"github.com/google/wire"
	"todolist-facebook-chatbot/conf"
)

func injectTaskRepository() (taskRepository, error) {
	wire.Build(conf.NewAppConfig, NewResource,
		wire.Struct(new(taskRepository), "resource", "appConfig"))
	return taskRepository{}, nil
}
