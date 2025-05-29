## 📘 `hautechai_storage` Resource

This resource manages storage records in the Hautech API.

### Attributes:

| Name       | Type                | Description |
|------------|---------------------|-------------|
| id         | string (computed)   | Unique identifier for the storage record. |
| key        | string (required)   | The key of the storage record. This is used to identify the record and must be unique. |
| value      | map (required)      | The value of the storage record as an object. |
| creator_id | string (computed)   | The ID of the creator of the storage record. |
| metadata   | string (computed)   | The metadata of the storage record as a JSON string. |
| created_at | string (computed)   | The creation timestamp of the storage record. |
| updated_at | string (computed)   | The last update timestamp of the storage record. |

### ⚠️ Behavior

The `key` attribute is immutable and requires replacement if changed.

On `create`:
- Creates a new storage record with the specified key and value.

On `update`:
- Updates the value of the existing storage record.

On `delete`:
- Deletes the storage record.

### Example

```hcl
resource "hautechai_storage" "example" {
  key   = "my-storage-key"
  value = {
    message = "Hello, World!"
    count   = "42"
    enabled = "true"
  }
}
```

### Example with Dynamic Value

```hcl
resource "hautechai_storage" "user_preferences" {
  key   = "user-${var.user_id}-preferences"
  value = {
    theme         = var.theme
    language      = var.language
    notifications = tostring(var.notifications_enabled)
  }
}
```

### Accessing Storage Data

```hcl
resource "hautechai_storage" "config" {
  key   = "app-config"
  value = {
    api_url     = "https://api.example.com"
    max_retries = "3"
    timeout     = "30"
  }
}

# Use the storage data in another resource
resource "hautechai_resource_access" "example" {
  resource_id    = hautechai_storage.config.value["api_url"]
  principal_id   = "account-456"
  principal_type = "account"
  access         = "reader"
}
```
