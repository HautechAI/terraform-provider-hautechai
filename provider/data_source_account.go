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

type AccountDataSource struct {
	client *hautechapi.ClientWithResponses
}

type AccountModel struct {
	Id types.String `tfsdk:"id"`
}

func NewAccountDataSource() datasource.DataSource {
	return &AccountDataSource{}
}

func (d *AccountDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account"
}

func (d *AccountDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches account",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (d *AccountDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData != nil {
		d.client = req.ProviderData.(*ProviderContext).Client
	}
}

func (d *AccountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AccountModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.Id.ValueString()
	balance, err := d.getAccountByID(ctx, accountId)
	if err != nil {
		resp.Diagnostics.AddError("Get Account Error", err.Error())
		return
	}

	data.Id = types.StringValue(balance.Id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccountDataSource) getAccountByID(ctx context.Context, id string) (*hautechapi.AccountEntity, error) {
	resp, err := r.client.AccountsControllerGetAccountV1WithResponse(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("account not found")
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return resp.JSON200, nil
}
