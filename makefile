include .env

create_migration:
	@read -p "Enter Migration Name: " migration; \
	migrate create -ext=sql -dir=internal/database/migrations -seq $$migration

migrate_up:
	migrate -path=internal/database/migrations -database "${DATABASE_URL}?sslmode=disable" -verbose up

migrate_down:
	migrate -path=internal/database/migrations -database "${DATABASE_URL}?sslmode=disable" -verbose down

migrate_force:
	@read -p "Enter Which Migration To Force To(Number): " migration; \
	migrate -path=internal/database/migrations -database "${DATABASE_URL}?sslmode=disable" force $$migration -verbose down

.PHONY: create_migration migrate_up migrate_down migrate_force