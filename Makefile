# create postgress container
postgres:
	sudo docker run --name postgres-12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres
createdb:
	sudo docker exec -it postgres-12 createdb --username=root --owner=root simple_bank
dropdb:
	sudo docker exec -it postgres-12 dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down -all
sqlc:
	sqlc generate
test:
	# verbose test with coverage details
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc
