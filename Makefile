DB_URL=postgresql://root:password@postgres:5432/postgres?sslmode=disable

network:
	docker network create bank-network

postgres:
	docker run --name postgres --network fleet-net -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=4545Moose -d postgres:14-alpine

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate
	
test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

.PHONY: network postgres migrateup migratedown sqlc test server mock 
