package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"todolist-facebook-chatbot/controllers"
	"todolist-facebook-chatbot/middlewares"
	"todolist-facebook-chatbot/validator"
)

func main() {
	gin := GetEngine()
	logrus.Info("Starting server...")
	logrus.Fatal(gin.Run(":8181"))
}

// Setup router
func GetEngine() *gin.Engine {
	app := gin.Default()
	binding.Validator = new(validator.DefaultValidator)

	app.Use(middlewares.NewGzip())
	app.Use(gin.Logger())
	app.Use(middlewares.NewCors([]string{"*"}))
	app.Use(middlewares.NewRecovery())
	app.GET("/swagger/*any", middlewares.NewSwagger())

	publicRoutes := app.Group("/v1")
	publicRoutes.Static("/public", "public")
	publicRoutes.Static("/tmp", "tmp")

	controllers.ApplyTaskAPI(publicRoutes)
	return app
}