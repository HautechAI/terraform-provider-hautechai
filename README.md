## Terraform provider for Hautech API
You can check api [here](https://api.hautech.ai/swagger)

## 🔧 How to use

### 1. Add the provider

```hcl
terraform {
  required_providers {
    hautech = {
      source  = "hautech/api"
      version = "0.1.0"
    }
  }
}
```

### 2. Configure the provider

```hcl
provider "hautech" {
  api_token = "your_token_here" # Replace with your real token
}
```

### 3. Use data source: fetch available permissions

```hcl
data "hautech_available_permissions" "permissions" {}
```

### 4. Output the result

```hcl
output "permissions" {
  value = data.hautech_available_permissions.permissions.items
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
  ~/.terraform.d/plugins/registry.terraform.io/hautech/api/0.1.0-dev/<os>_<arch>/terraform-provider-api
  ```

2. Then use it like a normal provider in your Terraform project:

```hcl
terraform {
  required_providers {
    hautech = {
      source  = "hautech/api"
      version = "0.1.0-dev"
    }
  }
}
```

No need to mess with `.terraformrc` or overrides — it just works locally ✨

## Docs

### 📘 `hautech_available_permissions` Data Source

This data source fetches the list of available permissions from the Hautech API.

---

### 🧱 Example Usage

```hcl
data "hautech_available_permissions" "permissions" {}
```

---

### 🔁 Output Attributes

| Name   | Type          | Description                              |
|--------|---------------|------------------------------------------|
| items  | `list(string)`| List of available permission identifiers |

---

### 🧾 Full Example

```hcl
provider "hautech" {
  api_token = "your_token_here"
}

data "hautech_available_permissions" "permissions" {}

output "permission_list" {
  value = data.hautech_available_permissions.permissions.items
}
```