sqlc:
	sqlc generate


migratecreate:
	migrate create -ext sql -dir db/schema/ -seq init_schema

migrateup:
	migrate -path db/schema -database "postgresql://postgres:postmoskwadb@uxti.de:5432/userdb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/schema -database "postgresql://postgres:postmoskwadb@uxti.de:5432/userdb?sslmode=disable" -verbose down

test:
	go test -v -cover ./...


.PHONY: protos

protos:
	protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto/user

clean:
	rm proto/user/*.go

run:
	go run main.go

.PHONY: sqlc migrateup migratedown
