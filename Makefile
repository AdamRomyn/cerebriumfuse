# Declare variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=main
BINARY_DIR=bin
CMD_DIR=./cmd/app

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v $(CMD_DIR)

# Test the project
test:
	$(GOTEST) -v ./...

# Clean up the binaries
clean:
	$(GOCLEAN)
	rm -f $(BINARY_DIR)/$(BINARY_NAME)

# Run the program
run:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v $(CMD_DIR)
	sudo ./$(BINARY_DIR)/$(BINARY_NAME) "./test_folders/nfs" "./test_folders/ssd" "/mnt/all-projects"


run_in_linux:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v $(CMD_DIR)
	sudo ./$(BINARY_DIR)/$(BINARY_NAME) "/mnt/shared_folder/cerebrium/test_folders/nfs" "/mnt/shared_folder/cerebrium/test_folders/ssd" "/mnt/all-projects"

setup:
	mkdir -p /mnt/all-projects
	rmdir ../ssd && mkdir ../ssd
	echo "File system setup for testing"