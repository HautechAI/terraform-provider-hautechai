## 📚 `hautechai_available_permissions` Data Source

This data source fetches the list of available permissions from the Hautech API.

### Example

```hcl
data "hautechai_available_permissions" "permissions" {}
```

### Output Attributes

| Name   | Type          | Description                              |
|--------|---------------|------------------------------------------|
| items  | `list(string)`| List of available permission identifiers |