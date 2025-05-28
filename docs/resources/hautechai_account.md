## 📘 `hautechai_account` Resource

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
resource "hautechai_account" "main" {
  alias = "example-account"

  group = {
    group_id = "group-123"
    role     = "maintainer"
  }
}
```