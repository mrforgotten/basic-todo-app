include .env
export

# PG_connection := "postgresql://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_NAME)?sslmode=disable"
db_connection:
	@echo "PostgreSQL Connection String: postgresql://$(PG_USERNAME):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DATABASE)"

migrateup:
	migrate -path ./db/migration -database "$(PG_connection)" -verbose up

migratedown:
	migrate -path ./db/migration -database "$(PG_connection)" -verbose down

migratecreate:
	migrate create -ext sql -dir ./db/migration -seq $(name)