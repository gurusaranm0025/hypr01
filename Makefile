BINARY_NAME=hyprone

INSTALL_DIR=/usr/local/bin
SCRIPTS_DIR=~/.local/share/bin

GO=go

all: build install

build:
	@echo "BUILDING THE BINARY..."
	@mkdir -p ./build
	$(GO) build -o ./build/$(BINARY_NAME)

install: build
	@echo "INSTALLING THE BINARY..."
	killall -9 hyprone
	sudo cp ./build/$(BINARY_NAME) $(INSTALL_DIR)
	sudo chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	
	@echo "PLACING SCRIPTS IN THEIR RIGHT PLACE..."
	cp ./scripts/* $(SCRIPTS_DIR)
	sudo chmod +x ~/.local/share/bin/*.sh
	
	@echo "STARTING HYPRONE..."
	hyprone -i & disown

clean:
	@echo "CLEANING UP..."
	rm -f ./build/$(BINARY_NAME)

run: build
	@echo "RUNNING THE BINARY..."
	./$(BINARY_NAME)

.PHONY: all build install clean run