#!/usr/bin/env bash

# Run integration tests
docker-compose -f deployment/docker-compose.yml -f deployment/integration-tests.yml \
    run newman-integration-tests
