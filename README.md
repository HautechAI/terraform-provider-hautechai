## Terraform provider for Hautech API
You can check API [here](https://api.hautech.ai/swagger)

## 🔧 How to use

### 1. Add the provider

```hcl
terraform {
  required_providers {
    hautechai = {
      source  = "HautechAI/hautechai"
      version = "${VERSION}"
    }
  }
}
```

### 2. Configure the provider

```hcl
provider "hautechai" {
  api_token = "your_token_here" # Replace with your real token
  api_url   = "https://custom-api.example.com" # Optional: Override the default API URL
}
```

#### Provider Configuration Parameters

| Name      | Type   | Required | Description                                                |
|-----------|--------|----------|------------------------------------------------------------|
| api_token | string | Yes      | Your Hautech API token                                     |
| api_url   | string | No       | Custom API URL. If not provided, the default URL will be used |

### 3. Use data source: fetch available permissions

```hcl
data "hautechai_available_permissions" "permissions" {}
```

### 4. Output the result

```hcl
output "permissions" {
  value = data.hautechai_available_permissions.permissions.items
}
```
---

## 🧪 Local development

To test the provider locally without publishing it to the Terraform Registry:

1. Run the helper script:

```bash
./scripts/build-for-local.sh
```

This will:
- Build the provider binary
- Copy it into your local Terraform plugin directory under:
  ```
  ~/.terraform.d/plugins/registry.terraform.io/HautechAI/hautechai/0.1.0-dev/<os>_<arch>/terraform-provider-api
  ```

2. Then use it like a normal provider in your Terraform project:

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