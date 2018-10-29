all: build test integration-test

bin:
	mkdir bin

build: bin
	go build -o bin/updates-notifier ./cmd/updates-notifier

integration-test: build
	./bin/updates-notifier

test:
	go test ./...

install:
	go install ./...
