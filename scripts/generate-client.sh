#!/usr/bin/env bash

docker run --rm -v "${PWD}/api:/local" openapitools/openapi-generator-cli:v5.0.0 generate \
    -i /local/doc.yml \
    -g php \
    -o /local/clients/go \
    -c /local/config-php.json

echo "Removing unused files"
rm -rfv api/clients/go/.openapi-generator api/clients/go/api

echo "Generation done!"
