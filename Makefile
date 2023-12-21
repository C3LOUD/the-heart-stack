MIGRATION_FILES_PATH="./db/migrations"

build:
		@go build -o ./bin ./cmd/the-heart-stack
migrate:
		@goose -dir $(MIGRATION_FILES_PATH) sqlite3 ./dev.db create $(filter-out $@,$(MAKECMDGOALS)) sql 
generate:
		@sqlc generate
up:
		@goose -dir $(MIGRATION_FILES_PATH) sqlite3 ./dev.db up
%:
		@: