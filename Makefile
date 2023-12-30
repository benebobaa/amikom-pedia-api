postgres:
	docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root amikom_pedia

dropdb:
	docker exec -it postgres16 dropdb amikom_pedia

migratecreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/amikom_pedia?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/amikom_pedia?sslmode=disable" -verbose down

PHONY: postgres createdb dropdb migrate_create migrate_up migrate_down