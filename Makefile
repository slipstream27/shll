BIN_DIR=bin
BIN_NAME=shll

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_NAME) ./cmd/shll

clean:
	rm -rf $(BIN_DIR)
