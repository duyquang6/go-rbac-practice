package main

import (
	"context"

	"github.com/duyquang6/go-rbac-practice/internal/api"
	"github.com/duyquang6/go-rbac-practice/internal/buildinfo"
	"github.com/duyquang6/go-rbac-practice/internal/database"
	"github.com/duyquang6/go-rbac-practice/internal/setup"
	"github.com/duyquang6/go-rbac-practice/pkg/logging"
	"github.com/sethvargo/go-signalcontext"
)

// main wrap realMain around a graceful shutdown scheme
func main() {
	ctx, done := signalcontext.OnInterrupt()

	logger := logging.NewLoggerFromEnv().
		With("build_id", buildinfo.RBACServer.ID()).
		With("build_time", buildinfo.RBACServer.Time())
	ctx = logging.WithLogger(ctx, logger)

	defer func() {
		done()
		if r := recover(); r != nil {
			logger.Fatalw("application panic", "panic", r)
		}
	}()

	err := realMain(ctx)
	done()

	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("successful shutdown")
}

func realMain(ctx context.Context) error {
	logger := logging.FromContext(ctx)

	var config database.Config
	env, err := setup.Setup(ctx, &config)
	if err != nil {
		logger.Fatal(err)
	}

	httpapp := api.NewHTTPServer(logger, env.Database())
	return httpapp.Run(ctx)
}
