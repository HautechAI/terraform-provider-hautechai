## Terraform provider for Hautech API
You can check API [here](https://api.hautech.ai/swagger)

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

---

## 📘 `hautech_account` Resource

This resource manages an account in the Hautech API.

### Attributes:

| Name        | Type          | Description |
|-------------|---------------|-------------|
| id          | string (computed) | Unique identifier returned by the API. Saved in state. |
| account_id  | string (optional) | ID to look up an existing account. If found, it will be used. If not found, the apply will fail. |
| alias       | string (optional) | Alias to find or create the account. If `account_id` is not set, it will use alias. |

### ⚠️ Behavior

This is an **immutable resource**. You cannot update existing accounts via Terraform.

On `create` and `update`, the logic is:

1. If `account_id` is provided:
  - Try to fetch the account by ID.
  - If not found, it throws an error.

2. Else if `alias` is provided:
  - Try to fetch the account by alias.
  - If not found, it creates a new account with that alias.

3. Else:
  - Create a brand new account with no alias.

### Example

```hcl
resource "hautech_account" "main" {
  alias = "example-account"
}
```

---

## 📘 `hautech_available_permissions` Data Source

This data source fetches the list of available permissions from the Hautech API.

### Example

```hcl
data "hautech_available_permissions" "permissions" {}
```

### Output Attributes

| Name   | Type          | Description                              |
|--------|---------------|------------------------------------------|
| items  | `list(string)`| List of available permission identifiers |

### Full Example

```hcl
provider "hautech" {
  api_token = "your_token_here"
}

data "hautech_available_permissions" "permissions" {}

output "permission_list" {
  value = data.hautech_available_permissions.permissions.items
}
```

---

## 📚 `hautech_account_balance` Data Source

This data source fetches the balance of a specific account by its `account_id`.

#### ⚙️ Example Usage

```hcl
data "hautech_account_balance" "balance" {
  account_id = "your-account-id-here"
}

output "balance" {
  value = data.hautech_account_balance.balance.amount
}
```

#### 📅 Input Parameters

| Name        | Type    | Required | Description                    |
|-------------|---------|----------|--------------------------------|
| account_id  | string  | ✅ Yes   | The ID of the account to fetch balance for |

#### 📄 Output Attributes

| Name    | Type   | Description         |
|---------|--------|---------------------|
| amount  | string | The current balance |

