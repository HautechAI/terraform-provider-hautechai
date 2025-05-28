# Hautech Terraform Provider Development Guidelines

This document provides guidelines and instructions for developing and maintaining the Hautech Terraform Provider.

## Build/Configuration Instructions

### Local Development Setup

1. **Prerequisites**:
   - Go 1.16 or later
   - Terraform CLI
   - Access to the Hautech API

2. **Building the Provider Locally**:
   
   Use the provided script to build and install the provider locally:
   ```bash
   ./scripts/build-for-local.sh
   ```
   
   This script:
   - Builds the provider binary
   - Copies it to your local Terraform plugin directory at:
     ```
     ~/.terraform.d/plugins/registry.terraform.io/HautechAI/hautechai/0.1.0-dev/<os>_<arch>/terraform-provider-hautechai
     ```

3. **Using the Local Provider**:
   
   In your Terraform configuration, specify the local version:
   ```hcl
   terraform {
     required_providers {
       hautechai = {
         source  = "HautechAI/hautechai"
         version = "0.1.0-dev"
       }
     }
   }
   ```

4. **API Code Generation**:
   
   If the Hautech API changes, regenerate the API client code:
   ```bash
   ./scripts/generate-api.sh
   ```

## Testing Information

### Running Tests

1. **Unit Tests**:
   
   Run all tests:
   ```bash
   go test -v ./...
   ```
   
   Run tests in a specific package:
   ```bash
   go test -v ./provider
   ```

2. **Test Dependencies**:
   
   If you add new dependencies, update the go.mod file:
   ```bash
   go mod tidy
   ```

### Adding New Tests

1. **Test File Naming**:
   
   Name test files with the `_test.go` suffix, in the same package as the code being tested.

2. **Test Function Naming**:
   
   Name test functions with the `Test` prefix followed by the name of the function or feature being tested.

3. **Example Test**:
   
   Here's a simple test for the provider:
   ```go
   package provider

   import (
       "testing"
       "github.com/stretchr/testify/assert"
   )

   // TestNewProvider tests that the provider factory function returns a non-nil provider
   func TestNewProvider(t *testing.T) {
       // Create a new provider using the factory function
       providerFactory := New(Config{
           ApiUrl: "https://api.hautech.ai",
       })
       
       // Call the factory function to get the provider
       p := providerFactory()
       
       // Check that the provider is not nil
       assert.NotNil(t, p, "Provider should not be nil")
   }
   ```

4. **Running the Example Test**:
   ```bash
   go test -v ./provider -run TestNewProvider
   ```

## Additional Development Information

### Code Style

1. **Formatting**:
   
   Use `gofmt` or `go fmt` to format your code:
   ```bash
   go fmt ./...
   ```

2. **Linting**:
   
   Use `golint` to check for style issues:
   ```bash
   golint ./...
   ```

3. **Error Handling**:
   
   Always check errors and return them with context:
   ```go
   if err != nil {
       return fmt.Errorf("failed to call API: %w", err)
   }
   ```

### Project Structure

1. **Provider Package**:
   
   The `provider` package contains the Terraform provider implementation, including:
   - Provider configuration (`provider.go`)
   - Resource implementations (`resource_*.go`)
   - Data source implementations (`data_source_*.go`)

2. **API Package**:
   
   The `api` package contains the generated API client code.

3. **Documentation**:
   
   Documentation for resources and data sources is in the `docs` directory.

### Adding New Resources or Data Sources

1. **Implementation**:
   
   Create a new file in the `provider` package for your resource or data source.

2. **Registration**:
   
   Register your resource or data source in the provider's `Resources()` or `DataSources()` method in `provider.go`.

3. **Documentation**:
   
   Add documentation for your resource or data source in the `docs` directory.

### Debugging

1. **Terraform Logs**:
   
   Enable Terraform logs for debugging:
   ```bash
   export TF_LOG=DEBUG
   ```

2. **Provider Logs**:
   
   Use the `tflog` package for logging in the provider:
   ```go
   import "github.com/hashicorp/terraform-plugin-log/tflog"

   tflog.Debug(ctx, "This is a debug message")
   ```