## 📘 `hautechai_resource_access` Resource

This resource grants access to a resource for a principal (account, user, group) in the Hautech API.

### Attributes:

| Name          | Type                | Description |
|---------------|---------------------|-------------|
| id            | string (computed)   | Unique identifier for the resource access grant. Generated from the combination of resource_id, principal_type, principal_id, and access. |
| resource_id   | string (required)   | The ID of the resource to grant access to. |
| principal_id  | string (required)   | The ID of the principal (account, user, group) to grant access to. |
| principal_type| string (required)   | The type of principal. Valid values: "account", "user", "group". |
| access        | string (required)   | The access level to grant. Valid values: "maintainer", "writer", "reader". |

### ⚠️ Behavior

This is an **immutable resource**. All attributes require replacement, meaning if any attribute changes, the resource will be destroyed and recreated.

On `create`:
- Grants the specified access to the resource for the principal.

On `delete`:
- Revokes the access from the resource for the principal.

### Example

```hcl
resource "hautechai_resource_access" "example" {
  resource_id    = "resource-123"
  principal_id   = "account-456"
  principal_type = "account"
  access         = "reader"
}
```

### Example with Account Resource

```hcl
resource "hautechai_account" "example" {
  alias = "example-account"
}

resource "hautechai_resource_access" "example_access" {
  resource_id    = "resource-123"
  principal_id   = hautechai_account.example.id
  principal_type = "account"
  access         = "reader"
}
```
