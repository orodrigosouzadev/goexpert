createmigration:
	@echo "Creating migration..."
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	@echo "Migrating..."
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

migratedown:
	@echo "Migrating down..."
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down

.PHONY: migrate migratedown createmigration