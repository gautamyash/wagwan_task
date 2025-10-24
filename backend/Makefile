.PHONY: run build clean

run:
	go run main.go

build:
	go build -o event-guest-manager main.go

clean:
	rm -f event-guest-manager

install-deps:
	go mod download
	go mod tidy

