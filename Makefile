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
test-cache:
	cd /mnt
	cat /mnt/all-projects/project-1/common-lib.py
	cat /mnt/all-projects/project-2/common-lib.py

	cat /mnt/all-projects/project-1/test-cache1.txt
	cat /mnt/all-projects/project-1/test-cache2.txt

	find ./test_folders/ssd -type f -exec cat {} \;

# Clean up the binaries
clean:
	$(GOCLEAN)
	rm -f $(BINARY_DIR)/$(BINARY_NAME)

unmount:
	umount /mnt/all-projects

# Run the program
run:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v $(CMD_DIR)
	./$(BINARY_DIR)/$(BINARY_NAME) "./test_folders/nfs" "./test_folders/ssd" "/mnt/all-projects"

setup:
	mkdir -p /mnt/all-projects
	rm -rf ./test_folders/ssd && mkdir ./test_folders/ssd
	echo "File system setup for testing"