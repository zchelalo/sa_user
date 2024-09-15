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

sqlc:
	sqlc generate

protouser:
	protoc --experimental_allow_proto3_optional \
	  --go_out=./pkg/proto --go_opt=paths=source_relative \
	  --go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative \
	  ./sa_proto/user/service.proto && \
	mv ./pkg/proto/sa_proto/user/* ./pkg/proto/user/ && \
	rm -rf ./pkg/proto/sa_proto

.PHONY: migrateup migrateup1 migratedown migratedown1 compose composebuild sqlc protousers
