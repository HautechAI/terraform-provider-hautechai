#!/bin/bash

SWAGGER_URL="https://api.hautech.ai/swagger.json"

brew install openapi-generator
rm -r ./api

openapi-generator generate \
  -i $SWAGGER_URL \
  -g go \
  -o ./api \
  --package-name hautechapi \
  -p generateInterfaces=true,isGoSubmodule=true,withGoMod=false \
  --global-property=apiTests=false,modelTests=false
