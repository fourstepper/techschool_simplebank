postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mokopass -d postgres:14-alpine

postgresrun:
	docker start postgres14

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:mokopass@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mokopass@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go simple-bank/db/sqlc Store

testcoverhtml:
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

.PHONY: postgres postgresrun createdb dropdb migrateup migratedown sqlc server mock testcoverhtml