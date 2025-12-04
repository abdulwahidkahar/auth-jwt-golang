run:
	go run cmd/api/main.go

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

migrate-create:
	@read -p "Enter migration name: " name; \
	touch db/migrations/$$(date +%s)_$$name.up.sql; \
	touch db/migrations/$$(date +%s)_$$name.down.sql
