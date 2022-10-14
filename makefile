migrateup: 
	migrate -path db/migration -database "postgresql://postgres:mercy@localhost:5432/airbnb?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://postgres:mercy@localhost:5432/airbnb?sslmode=disable" -verbose down 

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:mercy@localhost:5432/airbnb?sslmode=disable" -verbose down 1

migrateup1: 
	migrate -path db/migration -database "postgresql://postgres:mercy@localhost:5432/airbnb?sslmode=disable" -verbose up 1 
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/mysticis/airbnb_mini/db/sqlc Store

.PHONY: migrateup migrateup1 migratedown migratedown1 sqlc test server mock