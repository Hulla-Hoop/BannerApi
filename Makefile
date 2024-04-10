run:
	@docker-compose up -d
	@sleep 4
	@go run ./cmd/main.go
migrate-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=test password=12345678 port=5432 sslmode=disable" goose -dir ./migration/ up 
migrate-down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=test password=12345678 port=5432 sslmode=disable" goose -dir ./migration/ down 