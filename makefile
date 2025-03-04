DB_URL=postgres://postgres:postgres@localhost:5432/personal_dashboard?sslmode=disable

migrate-new:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -database "$(DB_URL)" -path migrations up

migrate-down:
	migrate -database "$(DB_URL)" -path migrations down 1

migrate-reset:
	migrate -database "$(DB_URL)" -path migrations down
	migrate -database "$(DB_URL)" -path migrations up