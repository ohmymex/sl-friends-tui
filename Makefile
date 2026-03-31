APP_NAME    := sl-friends
VERSION     := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS     := -s -w -X main.version=$(VERSION)

.PHONY: build run test lint clean

build:
	go build -ldflags "$(LDFLAGS)" -o bin/$(APP_NAME) ./cmd/sl-friends/

run:
	go run ./cmd/sl-friends/ $(ARGS)

test:
	go test ./... -race -count=1

lint:
	golangci-lint run

clean:
	rm -rf bin/
