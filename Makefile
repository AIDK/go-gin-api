build:
	@echo "Building the project"
	@go build -o bin/go-gin-api
	@echo "Build complete"

run: build
	@echo "Running the project"
	@./bin/go-gin-api
