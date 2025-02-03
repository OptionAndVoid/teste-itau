BIN_DIR = bin

CMD_DIR = cmd

SRC_DIR = cmd/teste-itau

.PHONY: all
all: help 

.PHONY: run
run:
	@go run $(SRC_DIR)/main.go

.PHONY: runtls
runtls:
	@go run $(SRC_DIR)/main.go ./ssl_credentials/server.key ./ssl_credentials/server.crt

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

.PHONY: build
build: $(BIN_DIR)
	go build -o $(BIN_DIR) ./$(CMD_DIR)/...

.PHONY: help
help:
	@echo "Usage:"
	@echo "  make run       - Run the API without TLS."
	@echo "  make runtls    - Run the API with TLS. Requires server.key and server.crt files in ./ssl_credentials/."
	@echo "  make build     - Build the API and place the binary in the $(BIN_DIR) directory."
	@echo "  make help      - Display this help message."

