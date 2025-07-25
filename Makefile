BINARY_NAME=hyprone

INSTALL_DIR=/usr/local/bin

GO=go

all: build install

build:
	@echo "BUILDING THE BINARY..."
	@mkdir -p ./build
	$(GO) build -o ./build/$(BINARY_NAME)

install: build
	@echo "INSTALLING THE BINARY TO $(INSTALL_DIR)"
	killall -9 hyprone
	sudo cp ./build/$(BINARY_NAME) $(INSTALL_DIR)
	sudo chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	hyprone -i & disown

clean:
	@echo "CLEANING UP..."
	rm -f ./build/$(BINARY_NAME)

run: build
	@echo "RUNNING THE BINARY..."
	./$(BINARY_NAME)

.PHONY: all build install clean run