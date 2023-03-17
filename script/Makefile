SHELL=cmd.exe
UI_BINARY=broker_service_app.exe
BROKER_BINARY=broker_service_app
AUTH_BINARY=auth_service_app

## up: starts all containers in the background without forcing build
up:
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started!

## up_build: stops docker-compose (if running), builds all projects and starts docker-compose
up_build: build_broker build_auth
	@echo Stopping docker images (if running...)
	docker-compose down
	@echo Building (when required) and starting docker images...
	docker-compose up --build -d --force-recreate
	@echo Docker images built and started!

## down: stop docker-compose
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!

## build_broker: builds the broker binary as a .exe
build_broker:
	@echo Building broker binary...
	chdir ..\broker && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0&& go build -o ${BROKER_BINARY} ./cmd/api
	@echo Done!

## build_auth: builds the authentication binary as a .exe
build_auth:
	@echo Building auth binary...
	chdir ..\authentication && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0&& go build -o ${AUTH_BINARY} ./cmd/api
	@echo Done!
