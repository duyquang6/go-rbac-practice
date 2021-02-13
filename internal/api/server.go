package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	authorizationSvc "github.com/duyquang6/go-rbac-practice/internal/authorization"
	authorizedDatabase "github.com/duyquang6/go-rbac-practice/internal/authorization/database"
	authorizationCon "github.com/duyquang6/go-rbac-practice/internal/controller/authorization"
	userCon "github.com/duyquang6/go-rbac-practice/internal/controller/user"
	userSvc "github.com/duyquang6/go-rbac-practice/internal/user"
	userDatabase "github.com/duyquang6/go-rbac-practice/internal/user/database"

	"github.com/duyquang6/go-rbac-practice/internal/database"
	"github.com/duyquang6/go-rbac-practice/internal/middleware"
	"github.com/duyquang6/go-rbac-practice/pkg/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type httpServer struct {
	logger *zap.SugaredLogger
	db     *database.DB
}

func NewHTTPServer(logger *zap.SugaredLogger, db *database.DB) *httpServer {
	return &httpServer{
		logger: logger,
		db:     db,
	}
}

func (s *httpServer) Run(ctx context.Context) error {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	v1.Use(middleware.PopulateRequestID())
	v1.Use(middleware.PopulateLogger(logging.FromContext(ctx)))
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
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	return s.ServeHTTP(ctx, srv)
}

// ServeHTTP starts the server and blocks until the provided context is closed.
// When the provided context is closed, the server is gracefully stopped with a
// timeout of 5 seconds.
//
// Once a server has been stopped, it is NOT safe for reuse.
func (s *httpServer) ServeHTTP(ctx context.Context, srv *http.Server) error {
	logger := logging.FromContext(ctx)

	// Spawn a goroutine that listens for context closure. When the context is
	// closed, the server is stopped.
	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()

		logger.Debugf("server.Serve: context closed")
		shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		logger.Debugf("server.Serve: shutting down")
		if err := srv.Shutdown(shutdownCtx); err != nil {
			select {
			case errCh <- err:
			default:
			}
		}
	}()

	// Run the server. This will block until the provided context is closed.
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	logger.Debugf("server.Serve: serving stopped")

	// Return any errors that happened during shutdown.
	select {
	case err := <-errCh:
		return fmt.Errorf("failed to shutdown: %w", err)
	default:
		return nil
	}
}
