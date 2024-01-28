-include ./scripts/Makefile

# Docker identifier
DOCKERNAME="microservice-starter"

## generate:client: Generate a PHP client
generate\:client:
	$(START)
	./scripts/generate-client.sh
	$(END)

## start:mock: Starts exposing external dependencies as mocks
start\:mock:
	$(START)
	docker-compose -f deployment/docker-compose.yml up --build "mock"
	$(END)
