DB_URL=postgresql://root:123456@localhost:5432/book_library?sslmode=disable

postgres:
	docker run --name books_container  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:14

createdb:
	docker exec -it books_container createdb --username=root --owner=root book_library

dropdb:
	docker exec -it books_container dropdb book_library

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc_gen:
	sqlc generate

run:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown new_migration sqlc_gen test run
