#!/bin/bash
REPO="github.com/duyquang6/go-rbac-practice"
NOW=$(date +'%Y-%m-%d_%T')

go build -ldflags "-X $REPO/internal/buildinfo.buildID=`git rev-parse --short HEAD` -X $REPO/internal/buildinfo.buildTime=$NOW" -o bin/migrate cmd/migrate/main.go

DB_NAME="rbac-db" DB_USER="dev" DB_PASSWORD="dev" DB_SSLMODE="disable" DB_HOST="0.0.0.0" DB_PORT="5432" ./bin/migrate