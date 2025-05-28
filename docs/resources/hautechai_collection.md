## 📘 `hautechai_collection` Resource

Manages a collection in the Hautech API.

### Example Usage

```hcl
resource "hautechai_collection" "collection" {
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
  value = hautechai_collection.collection
}
```

### Attributes

| Name     | Type             | Required | Description                      |
|----------|------------------|----------|----------------------------------|
| id       | `string`         | Computed | ID of the collection             |
| metadata | `map(string)`    | Optional | Key-value pairs as metadata      |
| items    | `list(string)`   | Optional | Item IDs to attach to collection |

### Behavior
- If `metadata` or `items` change, it syncs the collection accordingly.
- On update, items not present in the plan are removed.
- The `id` is assigned after the collection is created.

This resource supports creation, reading, updating, and deletion of Hautech collections via API.