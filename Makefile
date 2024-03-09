-include .env

devup:
	docker compose -f dev.env.compose.yml up -d

devdown:
	docker compose -f dev.env.compose.yml down

run:
	go run cmd/main.go

swag:
	swag init -g internal/transport/http/main.go -o api/openapi
