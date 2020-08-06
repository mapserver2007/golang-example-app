#!/bin/bash

set -xuC

pushd `dirname $0` > /dev/null

cd ..

rm -rf out

OPENAPI_GENERATOR_VERSION='4.3.1'
OPENAPI_GENERATE_LANGUAGE="go-server"
PROJECT_ROOT="${GOPATH}/src/github.com/mapserver2007/golang-example-app/web/openapi"
GIT_USER_ID="mapserver2007"
GIT_REPO_ID="golang-example-app/web/openapi/out"

docker run --rm -v ${PROJECT_ROOT}:/app \
  openapitools/openapi-generator-cli:v${OPENAPI_GENERATOR_VERSION} generate \
  -g ${OPENAPI_GENERATE_LANGUAGE} \
  --additional-properties=withInterfaces=true \
  -i /app/openapi-schema.yml \
  -o /app/out

popd
