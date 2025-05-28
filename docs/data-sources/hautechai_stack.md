## 📚 `hautechai_stack` Data Source

This data source fetches the specific stack by its `id`.

### ⚙️ Example Usage

```hcl
data "hautechai_stack" "stack" {
  id = "your-stack-id-here"
}

output "stack" {
  value = data.hautechai_stack.stack.id
}
```

### 📅 Input Parameters

| Name | Type    | Required | Description                    |
|------|---------|----------|--------------------------------|
| id   | string  | ✅ Yes   | The ID of the stack |