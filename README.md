## Terraform provider for Hautech API
You can check API [here](https://api.hautech.ai/swagger)

## 🔧 How to use

### 1. Add the provider

```hcl
terraform {
  required_providers {
    hautech = {
      source  = "hautech/api"
      version = "<!-- x-release-please-version -->"
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
| alias       | string (optional) | Alias to find or create the account. |
| group       | object (optional) | Optional group assignment (group_id + role). |

### Group Object Structure

```hcl
group = {
  group_id = "<group-id>"
  role     = "member" # Or: maintainer, owner
}
```

### ⚠️ Behavior

This is an **immutable resource**. You cannot update existing accounts via Terraform.

On `create` and `update`, the logic is:

1. If `alias` is provided:
- Try to fetch the account by alias.
- If not found, it creates a new account with that alias.

2. Else:
- Create a brand new account with no alias.

3. If `group` is set:
- Assign account to the given group with the specified role.
- If the group changes later, the old group is removed and the new one is added.

### Example

```hcl
resource "hautech_account" "main" {
  alias = "example-account"

  group = {
    group_id = "group-123"
    role     = "maintainer"
  }
}
```

## 📘 `hautech_collection` Resource

Manages a collection in the Hautech API.

#### Example Usage

```hcl
resource "hautech_collection" "collection" {
  metadata = {
    foo  = "bar"
    test = "123"
  }

  items = [
    "item-id-1",
    "item-id-2"
  ]
}

output "new_collection" {
  value = hautech_collection.collection
}
```

#### Attributes

| Name     | Type             | Required | Description                      |
|----------|------------------|----------|----------------------------------|
| id       | `string`         | Computed | ID of the collection             |
| metadata | `map(string)`    | Optional | Key-value pairs as metadata      |
| items    | `list(string)`   | Optional | Item IDs to attach to collection |

#### Behavior
- If `metadata` or `items` change, it syncs the collection accordingly.
- On update, items not present in the plan are removed.
- The `id` is assigned after the collection is created.

This resource supports creation, reading, updating, and deletion of Hautech collections via API.

## 📘 `hautech_stack` Resource

Manages a stack in the Hautech API.

#### Example Usage

```hcl
resource "hautech_stack" "stack" {
  metadata = {
    foo  = "bar"
    test = "123"
  }

  items = [
    "item-id-1",
    "item-id-2"
  ]
}

output "new_stack" {
  value = hautech_stack.stack
}
```

#### Attributes

| Name     | Type             | Required | Description                      |
|----------|------------------|----------|----------------------------------|
| id       | `string`         | Computed | ID of the stack             |
| metadata | `map(string)`    | Optional | Key-value pairs as metadata      |
| items    | `list(string)`   | Optional | Item IDs to attach to stack |

#### Behavior
- If `metadata` or `items` change, it syncs the stack accordingly.
- On update, items not present in the plan are removed.
- The `id` is assigned after the stack is created.

This resource supports creation, reading, updating, and deletion of Hautech stacks via API.


## 📘 `hautech_group` Resource

Manages a group in the Hautech API.

### Example Usage

```hcl
resource "hautech_group" "group" {}

output "group_id" {
  value = hautech_group.group.id
}
```

### Attributes

| Name | Type     | Required | Description          |
|------|----------|----------|----------------------|
| id   | `string` | Computed | ID of the group      |

### Behavior
- This resource is immutable. Any change requires recreating the resource.
- The `id` is generated by the backend after the group is created.

The resource supports create, read, and delete operations for Hautech groups. No update method is available.

---

## 📚 `hautech_available_permissions` Data Source

This data source fetches the list of available permissions from the Hautech API.

### Example

```hcl
data "hautech_available_permissions" "permissions" {}
```

### Output Attributes

| Name   | Type          | Description                              |
|--------|---------------|------------------------------------------|
| items  | `list(string)`| List of available permission identifiers |


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


## 📚 `hautech_account` Data Source

This data source fetches the specific account by its `id`.

#### ⚙️ Example Usage

```hcl
data "hautech_account" "account" {
  id = "your-account-id-here"
}

output "account" {
  value = data.hautech_account.account.id
}
```

#### 📅 Input Parameters

| Name | Type    | Required | Description                    |
|------|---------|----------|--------------------------------|
| id   | string  | ✅ Yes   | The ID of the account |


## 📚 `hautech_collection` Data Source

This data source fetches the specific collection by its `id`.

#### ⚙️ Example Usage

```hcl
data "hautech_collection" "collection" {
  id = "your-collection-id-here"
}

output "collection" {
  value = data.hautech_collection.collection.id
}
```

#### 📅 Input Parameters

| Name | Type    | Required | Description                    |
|------|---------|----------|--------------------------------|
| id   | string  | ✅ Yes   | The ID of the collection |

## 📚 `hautech_stack` Data Source

This data source fetches the specific stack by its `id`.

#### ⚙️ Example Usage

```hcl
data "hautech_stack" "stack" {
  id = "your-stack-id-here"
}

output "stack" {
  value = data.hautech_stack.stack.id
}
```

#### 📅 Input Parameters

| Name | Type    | Required | Description                    |
|------|---------|----------|--------------------------------|
| id   | string  | ✅ Yes   | The ID of the stack |