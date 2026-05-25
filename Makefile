run:
	go run cmd/api/main.go

worker:
	go run cmd/worker/main.go

test:
	go test ./...

swagger:
	swag init