## 📘 `hautechai_stack` Resource

Manages a stack in the Hautech API.

### Example Usage

```hcl
resource "hautechai_stack" "stack" {
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
  value = hautechai_stack.stack
}
```

### Attributes

| Name     | Type             | Required | Description                      |
|----------|------------------|----------|----------------------------------|
| id       | `string`         | Computed | ID of the stack                  |
| metadata | `map(string)`    | Optional | Key-value pairs as metadata      |
| items    | `list(string)`   | Optional | Item IDs to attach to stack      |

### Behavior
- If `metadata` or `items` change, it syncs the stack accordingly.
- On update, items not present in the plan are removed.
- The `id` is assigned after the stack is created.

This resource supports creation, reading, updating, and deletion of Hautech stacks via API.