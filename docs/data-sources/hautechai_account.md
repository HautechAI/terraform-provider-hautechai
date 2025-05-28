## 📚 `hautechai_account` Data Source

This data source fetches the specific account by its `id`.

### ⚙️ Example Usage

```hcl
data "hautechai_account" "account" {
  id = "your-account-id-here"
}

output "account" {
  value = data.hautechai_account.account.id
}
```

### 📅 Input Parameters

| Name | Type    | Required | Description                    |
|------|---------|----------|--------------------------------|
| id   | string  | ✅ Yes   | The ID of the account |