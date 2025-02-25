URI_DB=postgresql://postgres:example@localhost:5433/sa_user?sslmode=disable
MIGRATE=migrate -path pkg/sqlc/migration -database "$(URI_DB)" -verbose
DOCKER_COMPOSE_FILE = ./.dockers/docker-compose.yml

migrate-up:
	$(MIGRATE) up

migrate-up-1:
	$(MIGRATE) up 1

migrate-down:
	$(MIGRATE) down

migrate-down-1:
	$(MIGRATE) down 1

compose:
	docker compose -f $(DOCKER_COMPOSE_FILE) up

compose-build:
	docker compose -f $(DOCKER_COMPOSE_FILE) up --build

compose-build-detached:
	docker compose -f $(DOCKER_COMPOSE_FILE) up --build -d

create-envs:
	cp .env.example app.env

sqlc:
	sqlc generate

proto:
	protoc --experimental_allow_proto3_optional \
		-I=sa_proto/services \
		--go_out=./pkg/proto --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative \
		sa_proto/services/user.proto sa_proto/services/shared.proto

.PHONY: migrate-up migrate-up-1 migrate-down migrate-down-1 compose compose-build compose-build-detached create-envs sqlc proto
