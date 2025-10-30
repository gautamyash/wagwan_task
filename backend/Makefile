.PHONY: run build clean migrate-up migrate-down install-deps

run:
	go run main.go

build:
	go build -o event-guest-manager main.go

clean:
	rm -f event-guest-manager

install-deps:
	go mod download
	go mod tidy

migrate-up:
	go run ./cmd/migrate

migrate-down:
	go run ./cmd/migrate down
