package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	hautechapi "hautech/api"
	"net/http"
)

type StackDataSource struct {
	client *hautechapi.ClientWithResponses
}

type StackModel struct {
	ID       types.String `tfsdk:"id"`
	Metadata types.Map    `tfsdk:"metadata"`
}

func NewStackDataSource() datasource.DataSource {
	return &StackDataSource{}
}

func (d *StackDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_stack"
}

func (d *StackDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches Stack",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"metadata": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (d *StackDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData != nil {
		d.client = req.ProviderData.(*Context).Client
	}
}

func (d *StackDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data StackModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	stack, err := d.getStackByID(ctx, data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Get Stack Error", err.Error())
		return
	}

	data.ID = types.StringValue(stack.Id)
	data.Metadata, _ = types.MapValueFrom(ctx, types.StringType, stack.Metadata)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *StackDataSource) getStackByID(ctx context.Context, id string) (*hautechapi.StackEntity, error) {
	resp, err := d.client.StacksControllerGetStackV1WithResponse(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("stack not found")
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return resp.JSON200, nil
}
