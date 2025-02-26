include app.env
export $(shell sed 's/=.*//' app.env)

DOCKER_COMPOSE_FILE = ./.dockers/docker-compose.yml
DOCKER_NETWORK = dockers_sa_user_network

URI_DB := postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATE := docker run -v $(shell pwd)/pkg/sqlc/migration:/migrations --network $(DOCKER_NETWORK) migrate/migrate -path /migrations -database "$(URI_DB)" -verbose

setup:
	$(MAKE) create-envs
	$(MAKE) compose-build-detached
	docker run --rm --network=$(DOCKER_NETWORK) \
		-v $(shell pwd)/scripts:/scripts alpine sh /scripts/wait_for_db.sh $(DB_HOST) $(DB_PORT)
	$(MAKE) migrate-up

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

compose-down:
	docker compose -f $(DOCKER_COMPOSE_FILE) down

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

.PHONY: setup migrate-up migrate-up-1 migrate-down migrate-down-1 compose compose-build compose-build-detached compose-down create-envs sqlc proto
