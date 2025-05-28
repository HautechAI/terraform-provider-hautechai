## 📚 `hautechai_collection` Data Source

This data source fetches the specific collection by its `id`.

### ⚙️ Example Usage

```hcl
data "hautechai_collection" "collection" {
  id = "your-collection-id-here"
}

output "collection" {
  value = data.hautechai_collection.collection.id
}
```

### 📅 Input Parameters

| Name | Type    | Required | Description                    |
|------|---------|----------|--------------------------------|
| id   | string  | ✅ Yes   | The ID of the collection |