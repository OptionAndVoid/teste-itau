BIN_DIR = bin

CMD_DIR = cmd

SRC_DIR = cmd/teste-itau


.PHONY: all
all: run

.PHONY: run
run:
	@go run $(SRC_DIR)/main.go

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

.PHONY: build
build: $(BIN_DIR)
	go build -o $(BIN_DIR) ./$(CMD_DIR)/...
