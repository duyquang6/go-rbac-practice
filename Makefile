POSTGRESQL_URL:=http://google.com
migrate: ## Migrate DB
	@read -p "Enter migration name: " migration_name; \
	echo Executing migrate . . .; \
	migrate create -ext sql -dir database/migrations -seq $${migration_name}

migration: ## Migrate DB
	@migrate -database $(POSTGRESQL_URL) -path database/migrations up
