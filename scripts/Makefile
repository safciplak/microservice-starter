# DO NOT EDIT. This file can ONLY be modified from the docker-images repo and will be overwritten otherwise.

.DEFAULT_GOAL := help
.PHONY: upgrade generate generate\:mocks generate\:wire start watch release db\:initialize db\:migrate db\:fresh db\:seed test lint format bash help

START=@echo "$(ANOTHERCOLOR) > Starting $@ ...$(NOCOLOR)"
END=@echo "$(OUTPUTCOLOR) > Finished $@ ...$(NOCOLOR)"

# Colors used
NOCOLOR=\033[0m
INPUTCOLOR=\033[0;36m
OUTPUTCOLOR=\033[0;32m
SOMECOLOR=\033[01;31m
ANOTHERCOLOR=\033[01;34m

ARG1=""
ARG2=""

RUN_DEBUG=docker compose -f deployment/docker-compose.yml exec "${DOCKERNAME}-debug"
RUN_POSTGRES=docker compose -f deployment/docker-compose.yml exec postgres
START_DOCKER=docker compose -f deployment/docker-compose.yml up --build

# Checks if the DOCKERNAME variable is set correctly
check_valid_dockername:
	@if [[ $(DOCKERNAME) == *"-debug" ]]; then \
  		echo >&2 "Invalid DOCKERNAME found, please remove -debug from the Makefile DOCKERNAME and make sure\
  		to use double quotes."; exit 1; \
	fi

# This is triggered when the command was not found
%::
	make
	@echo "$(SOMECOLOR) > type one of the targets above$(NOCOLOR)"
	@echo

## upgrade: Upgrade all go dependencies
upgrade:
	$(START)
	go get -u all
	$(END)

## generate: Generates all go annotations
generate:
	$(START)
	@$(RUN_DEBUG) sh -c "PWD=/app go generate ./..."
	$(END)

## generate:mocks: Generate all mockfiles
generate\:mocks:
	$(START)
	@$(RUN_DEBUG) generate-mocks.sh
	$(END)

## generate:wire: Generate initial wire_gen file
generate\:wire:
	$(START)
	@$(RUN_DEBUG) wire
	$(END)

## start: Starts the application
start:
	$(START)
	$(START_DOCKER) "${DOCKERNAME}-build"
	$(END)

## watch: Starts the application with the Delve debugger
watch: check_valid_dockername
	$(START)
	$(START_DOCKER) "${DOCKERNAME}-debug"
	$(END)

## release: Starts the application in release mode
release:
	$(START)
	@$(START_DOCKER) "${DOCKERNAME}-release"
	$(END)

## db:initialize: Initializes the database
db\:initialize:
	$(START)
	@$(RUN_POSTGRES) /docker-entrypoint-initdb.d/initialize-database.sh
	$(END)

## db:migrate: Runs migrations
db\:migrate:
	$(START)
	@$(RUN_DEBUG) sh -c "cd src && capila migrate up"
	$(END)

## db:fresh: Fresh DB
db\:fresh:
	$(START)
	@$(RUN_DEBUG) sh -c "cd src && capila migrate fresh"
	$(END)

## db:seed: Seed DB
db\:seed:
	$(START)
	@$(RUN_POSTGRES) sh -c 'find /seeds -iname "*.up.sql" -exec psql -U $$POSTGRES_USER -d $$POSTGRES_DB -a -f {} \;'
	$(END)

## test: Runs unit tests
test:
	$(START)
	@$(RUN_DEBUG) go test ./... -count=1 -failfast
	$(END)

## lint: Runs linting utilities and uploads to SonarQube
lint:
	$(START)
	@$(RUN_DEBUG) golangci-lint run
	$(END)

## tidy: Runs go mod tidy command
tidy:
	$(START)
	@$(RUN_DEBUG) go mod tidy
	$(END)

## format: Runs gofmt on all code
format:
	$(START)
	@$(RUN_DEBUG) gofmt -s -w .
	$(END)

## bash: Go into the docker instance..
bash:
	$(START)
	@$(RUN_DEBUG) bash
	$(END)

makefile: help
help: Makefile
	@echo "$(ANOTHERCOLOR) > Choose a make command from the following:"
	@{ echo " -----------: -----------\n COMMAND: DESCRIPTION\n -----------: -----------$(NOCOLOR)" && \
	(fgrep -h "## " $(MAKEFILE_LIST) | sed -e 's/##//' | sort | sed -e '1d') } | sed -e 's/: /։/' | column -t -s \։
	@echo
