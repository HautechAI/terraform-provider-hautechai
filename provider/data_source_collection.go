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

type CollectionDataSource struct {
	client *hautechapi.ClientWithResponses
}

type CollectionModel struct {
	ID       types.String `tfsdk:"id"`
	Metadata types.Map    `tfsdk:"metadata"`
}

func NewCollectionDataSource() datasource.DataSource {
	return &CollectionDataSource{}
}

func (d *CollectionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_collection"
}

func (d *CollectionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches Collection",
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

func (d *CollectionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData != nil {
		d.client = req.ProviderData.(*Context).Client
	}
}

func (d *CollectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CollectionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	collection, err := d.getCollectionByID(ctx, data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Get Collection Error", err.Error())
		return
	}

	data.ID = types.StringValue(collection.Id)
	data.Metadata, _ = types.MapValueFrom(ctx, types.StringType, collection.Metadata)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *CollectionDataSource) getCollectionByID(ctx context.Context, id string) (*hautechapi.CollectionEntity, error) {
	resp, err := d.client.CollectionsControllerGetCollectionV1WithResponse(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("collection not found")
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return resp.JSON200, nil
}
