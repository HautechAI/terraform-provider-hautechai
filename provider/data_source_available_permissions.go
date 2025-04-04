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

type AvailablePermissionsDataSource struct {
	client *hautechapi.ClientWithResponses
}

type AvailablePermissionsModel struct {
	Items []types.String `tfsdk:"items"`
}

func NewAvailablePermissionsDataSource() datasource.DataSource {
	return &AvailablePermissionsDataSource{}
}

func (d *AvailablePermissionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_available_permissions"
}

func (d *AvailablePermissionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches available permissions from the Hautech API.",
		Attributes: map[string]schema.Attribute{
			"items": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
		},
	}
}

func (d *AvailablePermissionsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData != nil {
		d.client = req.ProviderData.(*Context).Client
	}
}

func (d *AvailablePermissionsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	permissions, err := d.getPermissions(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Get Permissions Error", err.Error())
		return
	}

	var tfItems []types.String
	for _, item := range *permissions {
		tfItems = append(tfItems, types.StringValue(item))
	}

	state := AvailablePermissionsModel{
		Items: tfItems,
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (d *AvailablePermissionsDataSource) getPermissions(ctx context.Context) (*[]string, error) {
	resp, err := d.client.PermissionsControllerListAvailablePermissionsV1WithResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("unexpected status: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return resp.JSON200, nil
}
