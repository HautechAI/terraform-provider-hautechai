package provider

import (
	"context"
	"fmt"
	"net/http"

	hautechapi "hautech/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountBalanceDataSource struct {
	client *hautechapi.ClientWithResponses
}

type AccountBalanceModel struct {
	AccountId types.String `tfsdk:"account_id"`
	Amount    types.String `tfsdk:"amount"`
}

func NewAccountBalanceDataSource() datasource.DataSource {
	return &AccountBalanceDataSource{}
}

func (d *AccountBalanceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account_balance"
}

func (d *AccountBalanceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches available account balance",
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"amount": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *AccountBalanceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData != nil {
		d.client = req.ProviderData.(*ProviderContext).Client
	}
}

func (d *AccountBalanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AccountBalanceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountId.ValueString()
	balance, err := d.getAccountBalance(ctx, accountId)
	if err != nil {
		resp.Diagnostics.AddError("Get Balance Error", err.Error())
		return
	}

	data.Amount = types.StringValue(balance.Balance)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *AccountBalanceDataSource) getAccountBalance(ctx context.Context, id string) (*hautechapi.BalanceResultDto, error) {
	resp, err := d.client.BalancesControllerGetBalanceV1WithResponse(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("balance not found")
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return resp.JSON200, nil
}
