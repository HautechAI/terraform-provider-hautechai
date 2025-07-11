## 📘 `hautechai_group` Resource

Manages a group in the Hautech API.

### Example Usage

```hcl
resource "hautechai_group" "group" {}

output "group_id" {
  value = hautechai_group.group.id
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