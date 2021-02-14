package api

import (
	"context"

	authSvc "github.com/duyquang6/go-rbac-practice/internal/auth"
	authorizationSvc "github.com/duyquang6/go-rbac-practice/internal/authorization"
	authorizedDatabase "github.com/duyquang6/go-rbac-practice/internal/authorization/database"
	authCon "github.com/duyquang6/go-rbac-practice/internal/controller/auth"
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
	v1 := r.Group("/v1")
	authorRepo := authorizedDatabase.New(s.db)
	authorService := authorizationSvc.NewAuthorizationService(authorRepo)
	authorController := authorizationCon.New(authorService)

	userRepo := userDatabase.New(s.db)
	userService := userSvc.New(userRepo)
	userController := userCon.New(userService)

	authService := authSvc.New(userRepo, authorRepo)
	authController := authCon.New(authService)

	populateSession := middleware.PopulateSessionIfNotExist(sessionStore)

	v1.Use(middleware.PopulateRequestID())
	v1.Use(middleware.PopulateLogger(logging.FromContext(ctx)))
	v1.Use(populateSession)
	{
		{
			v1.POST("/login", authController.HandleLogin())
		}

		sub := v1.Group("/")
		sub.Use(middleware.AuthSession(authService, userService))
		{
			// role
			sub.POST("/role", authorController.HandleCreateRole())
			sub.POST("/role/:id/binding", authorController.HandleBindingPolicyRole())

			// policy
			sub.POST("/policy", authorController.HandleCreatePolicy())
			sub.POST("/policy/:id/append", authorController.HandleAppendPermissionPolicy())
		}

		{
			// role
			sub.POST("/user", userController.HandleCreateUser())
			sub.POST("/user/:id/binding", userController.HandleBindingRoleUser())
		}

	}
}
