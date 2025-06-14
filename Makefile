HELLO_BETS_BINARY=helloBetsApp

up:
	@echo "Starting the application..."
	docker compose up -d

down:
	@echo "Stopping the application..."
	docker compose down -v

build:
	@echo "Building the application..."
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(HELLO_BETS_BINARY) ./cmd/
	@echo "Building the Docker image..."
	docker build -t hello-bets-app .
	@echo "Docker image built successfully."

build-all: build
	@echo "Building and starting the application..."
	docker compose build

br: build up
	@echo "Building and starting the application..."