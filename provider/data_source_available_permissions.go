package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	hautechapi "hautech/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &AvailablePermissionsDataSource{}

type AvailablePermissionsDataSource struct {
	client *hautechapi.ClientWithResponses
}

type availablePermissionsModel struct {
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
	if req.ProviderData == nil {
		return
	}

	ctx := req.ProviderData.(*ProviderContext)
	d.client = ctx.Client
}

func (d *AvailablePermissionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	apiResp, err := d.client.PermissionsControllerListAvailablePermissionsV1WithResponse(ctx)
	if err != nil {
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("failed to list available permissions: %v", err))
		return
	}

	if apiResp.StatusCode() != 200 {
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("unexpected status code: %d", apiResp.StatusCode()))
		return
	}

	permissions := apiResp.JSON200
	if permissions == nil {
		resp.Diagnostics.AddError("API Error", "received nil permissions from API")
		return
	}

	var tfItems []types.String
	for _, item := range *permissions {
		tfItems = append(tfItems, types.StringValue(item))
	}

	state := availablePermissionsModel{
		Items: tfItems,
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	tflog.Trace(ctx, "Retrieved available permissions", map[string]interface{}{
		"count": len(tfItems),
	})
}
