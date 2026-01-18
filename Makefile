BINARY_NAME=hyprone

INSTALL_DIR=/usr/local/bin
SCRIPTS_DIR=~/.local/share/bin

GO=go

all: build install

build:
	@echo "BUILDING BINARY..."
	@mkdir -p ./build
	$(GO) build -o ./build/$(BINARY_NAME)

install: build
	@echo "KILLING hyprone PROCESS..."
	-killall -9 hyprone

	@echo "INSTALLING..."
	sudo cp ./build/$(BINARY_NAME) $(INSTALL_DIR)
	sudo install -Dm755 ./build/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)

	@echo "PERFORMING INITIAL SETUP..."
	hyprone --initial-setup full

	@echo "STARTING hyprone..."
	hyprone -i & disown

clean:
	@echo "CLEANING UP..."
	rm -f ./build/$(BINARY_NAME)

run: build
	@echo "RUNNING THE BINARY..."
	./$(BINARY_NAME)

.PHONY: all build install clean run