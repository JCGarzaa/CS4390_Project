# Makefile

# Compiler and Compiler Flags
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_FOLDER=bin

# Target parameters
ARGS ?= "default_name"

# Targets
all: server client

server:
	@mkdir -p $(BINARY_FOLDER)
	$(GOBUILD) -o $(BINARY_FOLDER)/server ./Server/server.go

client:
	@mkdir -p $(BINARY_FOLDER)
	$(GOBUILD) -o $(BINARY_FOLDER)/client ./Client/client.go

clean:
	$(GOCLEAN)
	rm -f $(BINARY_FOLDER)/server
	rm -f $(BINARY_FOLDER)/client

run-server:
	./$(BINARY_FOLDER)/server

run-client:
	./$(BINARY_FOLDER)/client $(ARGS)

.PHONY: server client clean run-server run-client
