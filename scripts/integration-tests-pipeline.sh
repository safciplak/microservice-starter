#!/bin/bash
echo "Testing branch: $CI_BRANCH"

# Extra tags.
if [[ "$CI_BRANCH" =~ staging\/uat.* ]]; then
  newman run microservice-starter.postman_collection.json -e microservice-starter."uat${CI_BRANCH: -1}".postman_environment.json

elif [[ "$CI_BRANCH" =~ master ]] || [[ "$CI_BRANCH" =~ hotfix\/uat.* ]]; then
  newman run microservice-starter.postman_collection.json -e microservice-starter.live.postman_environment.json
else
  echo "Skipping the integration tests";
fi
