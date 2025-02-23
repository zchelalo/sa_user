URI_DB=postgresql://postgres:example@localhost:5433/sa_user?sslmode=disable
MIGRATE=migrate -path pkg/sqlc/migration -database "$(URI_DB)" -verbose

migrateup:
	$(MIGRATE) up

migrateup1:
	$(MIGRATE) up 1

migratedown:
	$(MIGRATE) down

migratedown1:
	$(MIGRATE) down 1

compose:
	docker compose -f ./.dockers/docker-compose.yml up

composebuild:
	docker compose -f ./.dockers/docker-compose.yml up --build

composebuilddetached:
	docker compose -f ./.dockers/docker-compose.yml up --build -d

sqlc:
	sqlc generate

proto:
	protoc --experimental_allow_proto3_optional \
		-I=sa_proto/services \
		--go_out=./pkg/proto --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative \
		sa_proto/services/user.proto sa_proto/services/shared.proto

.PHONY: migrateup migrateup1 migratedown migratedown1 compose composebuild sqlc proto
