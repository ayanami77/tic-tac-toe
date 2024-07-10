PROJECT_NAME=tic-tac-toe

build:
		@echo "Building the project..."
		@go build -o bin/$(PROJECT_NAME) ./main.go
		@echo "Build successful!"

test:
		@echo "Running tests..."
		@go test -v ./...

release-check:
		goreleaser check

release-snapshot:
		goreleaser release --snapshot --clean
