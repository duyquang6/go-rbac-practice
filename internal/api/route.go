package api

import (
	"context"

	authorizationSvc "github.com/duyquang6/go-rbac-practice/internal/authorization"
	authorizedDatabase "github.com/duyquang6/go-rbac-practice/internal/authorization/database"
	authorizationCon "github.com/duyquang6/go-rbac-practice/internal/controller/authorization"
	userCon "github.com/duyquang6/go-rbac-practice/internal/controller/user"
	"github.com/duyquang6/go-rbac-practice/internal/middleware"
	userSvc "github.com/duyquang6/go-rbac-practice/internal/user"
	userDatabase "github.com/duyquang6/go-rbac-practice/internal/user/database"
	"github.com/duyquang6/go-rbac-practice/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func (s *httpServer) initRoute(ctx context.Context, r *gin.Engine, sessionStore sessions.Store) {
	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	populateSession := middleware.PopulateSessionIfNotExist(sessionStore)
	v1 := r.Group("/v1")
	v1.Use(middleware.PopulateRequestID())
	v1.Use(middleware.PopulateLogger(logging.FromContext(ctx)))
	v1.Use(populateSession)
	{
		authorRepo := authorizedDatabase.New(s.db)
		authorService := authorizationSvc.NewAuthorizationService(authorRepo)
		authorController := authorizationCon.New(authorService)
		{
			// role
			v1.POST("/role", authorController.HandleCreateRole())
			v1.POST("/role/:id/binding", authorController.HandleBindingPolicyRole())

			// policy
			v1.POST("/policy", authorController.HandleCreatePolicy())
			v1.POST("/policy/:id/append", authorController.HandleAppendPermissionPolicy())
		}

		userRepo := userDatabase.New(s.db)
		userService := userSvc.New(userRepo)
		userController := userCon.New(userService)
		{
			// role
			v1.POST("/user", userController.HandleCreateUser())
			v1.POST("/user/:id/binding", userController.HandleBindingRoleUser())
		}
	}
}
