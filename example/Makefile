.PHONY: build generate clean

# Default target (build will be executed when you run `make`)
build: clean generate
	@echo "Tidying up Go modules..."
	@go mod tidy
	@echo "Building the application..."
	@go build -trimpath -gcflags=all='-l' -ldflags "-s -w" -o bin/app main.go

# Run code generation
generate:
	@echo "Running go generate..."
	@go generate ./...

# Clean up generated files and binaries
clean:
	@echo "Cleaning up..."
	@rm -rf bin/
