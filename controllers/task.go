package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist-facebook-chatbot/dtos"
	"todolist-facebook-chatbot/services"
)

type TaskController struct {
	taskService services.ITaskService
}

func ApplyTaskAPI(app *gin.RouterGroup) {
	c := &TaskController{}
	_ = services.GetServiceContainer().Invoke(func(taskService services.ITaskService) {
		c.taskService = taskService
	})
	taskRoutes := app.Group("/tasks")
	{
		taskRoutes.POST("", c.createTask())
	}
}

// Create inserts a new Task to database.
// @Title Task - Create
// @Description Inserts a new Task to database
// @Param body body dtos.CreateTaskRequest true "The task information for insert"
// @Param    Authorization    	header    	string    	true    	"Access Token."
// @Success 200 {object} dtos.CreateTaskResponse
// @Failure 400 Bad Request
// @Failure 500 Internal Server Error
// @router / [post]
func (s *TaskController) createTask() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		body := dtos.CreateTaskRequest{}
		if err := ctx.Bind(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := s.taskService.Create(&body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}
