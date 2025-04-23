#!/bin/bash

OS=$(go env GOOS)
ARCH="$(go env GOARCH)"
VERSION="0.1.0-dev"
PROVIDER_NAME="hautechai"
NAMESPACE="HautechAI"
REGISTRY="registry.terraform.io"
REGISTRY_PATH="$HOME/.terraform.d/plugins/$REGISTRY/$NAMESPACE/$PROVIDER_NAME/$VERSION/${OS}_${ARCH}"

mkdir -p "$REGISTRY_PATH"

go build -o "$REGISTRY_PATH/terraform-provider-$PROVIDER_NAME"

echo "✅ Provider built at: $REGISTRY_PATH/terraform-provider-$PROVIDER_NAME"