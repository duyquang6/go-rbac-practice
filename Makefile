POSTGRESQL_URL:=http://google.com
migration: ## Create new migration DB
	@read -p "Enter migration name: " migration_name; \
	echo Executing migrate . . .; \
	migrate create -ext sql -dir database/migrations -seq $${migration_name}

migrate: ## Migrate DB
	@migrate -database $(POSTGRESQL_URL) -path database/migrations up
