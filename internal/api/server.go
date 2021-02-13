package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/duyquang6/go-rbac-practice/internal/database"
	"github.com/duyquang6/go-rbac-practice/pkg/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/boj/redistore.v1"
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
	// Setup sessions
	sessionStore, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte(""))
	if err != nil {
		s.logger.Fatal("cannot init redis session: %v", err)
	}
	s.initRoute(ctx, r, sessionStore)
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
