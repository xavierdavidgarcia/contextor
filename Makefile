.PHONY: build install clean

# Build the binary
build:
	go build -o contextor && chmod +x contextor

# Install the binary to /usr/local/bin (requires sudo)
install-system: build
	sudo cp contextor /usr/local/bin/

# Install the binary to ~/bin (creates the directory if it doesn't exist)
install-user: build
	mkdir -p ~/bin
	cp contextor ~/bin/
	@echo ""
	@echo "Binary installed to ~/bin/contextor"
	@echo "If ~/bin is not in your PATH, add the following to your shell profile:"
	@echo "    export PATH=$$PATH:$$HOME/bin"
	@echo ""

# Default install (user install)
install: install-system

# Clean build artifacts
clean:
	rm -f contextor
