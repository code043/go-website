BIN_NAME = bin/website

MAIN = main.go

dev:
	@go run $(MAIN)

test:
	@go test -v ./...

build:
	@go build -o $(BIN_NAME) $(MAIN)

run: build
	@./$(BIN_NAME)

clean:
	@rm -rf bin

format:
	@go fmt ./...

lint:
	@golangci-lint run
