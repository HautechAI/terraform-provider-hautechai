package provider

import (
	"context"

	hautechapi "hautech/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &AvailablePermissionsDataSource{}

type AvailablePermissionsDataSource struct {
	client *hautechapi.APIClient
}

func NewAvailablePermissionsDataSource() datasource.DataSource {
	return &AvailablePermissionsDataSource{}
}

func (d *AvailablePermissionsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "hautech_available_permissions"
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

type availablePermissionsModel struct {
	Items []types.String `tfsdk:"items"`
}

func (d *AvailablePermissionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	apiResp, _, err := d.client.PermissionsAPI.PermissionsControllerListAvailablePermissionsV1(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError("API Error", err.Error())
		return
	}

	var tfItems []types.String
	for _, item := range apiResp {
		tfItems = append(tfItems, types.StringValue(item))
	}

	state := availablePermissionsModel{
		Items: tfItems,
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
