createDb:
	createdb --username=postgres --owner=postgres go_finance
dropDb:
	dropdb --username=postgres go_finance
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine
migrateUp:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up
migrateDrop:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose drop
sqlcGenerate:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
test:
	go test -v ./...
server:
	go run main.go

.PHONY: createDb dropDb postgres migrateUp migrateDown sqlcGenerate test server