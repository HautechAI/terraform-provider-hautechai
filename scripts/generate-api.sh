#!/bin/bash

SWAGGER_URL="https://api.hautech.ai/swagger.json"

go tool oapi-codegen --config=oapi-codegen.yml $SWAGGER_URL
