goose-status:
	go run ./cmd/goose/main.go status

goose-up:
	go run ./cmd/goose/main.go up

goose-down:
	go run ./cmd/goose/main.go down

run:
	go run ./cmd/api/main.go
