include .env
export

migrate-up:
	migrate -path database/migrations -database "mysql://${MYSQL_URI}" -verbose up

migrate-down:
	migrate -path database/migrations -database "mysql://${MYSQL_URI}" -verbose down
