## 📚 `hautechai_account_balance` Data Source

This data source fetches the balance of a specific account by its `account_id`.

### ⚙️ Example Usage

```hcl
data "hautechai_account_balance" "balance" {
  account_id = "your-account-id-here"
}

output "balance" {
  value = data.hautechai_account_balance.balance.amount
}
```

### 📅 Input Parameters

| Name        | Type    | Required | Description                    |
|-------------|---------|----------|--------------------------------|
| account_id  | string  | ✅ Yes   | The ID of the account to fetch balance for |

### 📄 Output Attributes

| Name    | Type   | Description         |
|---------|--------|---------------------|
| amount  | string | The current balance |