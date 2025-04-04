package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	hautechapi "hautech/api"
	"net/http"
)

type CollectionResource struct {
	client *hautechapi.ClientWithResponses
}

type CollectionResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Metadata types.Map    `tfsdk:"metadata"`
}

func NewCollectionResource() resource.Resource {
	return &CollectionResource{}
}

func (r *CollectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_collection"
}

func (r *CollectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage Hautech Collections.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"metadata": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *CollectionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData != nil {
		r.client = req.ProviderData.(*Context).Client
	}
}

func (r *CollectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CollectionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	metadata := make(map[string]interface{}, len(data.Metadata.Elements()))
	for k, v := range data.Metadata.Elements() {
		strVal, ok := v.(types.String)
		if !ok || strVal.IsNull() {
			continue
		}
		metadata[k] = strVal.ValueString()
	}

	collection, err := r.createCollection(ctx, hautechapi.CreateCollectionParamsDto{Metadata: &metadata})
	if err != nil {
		resp.Diagnostics.AddError("Create Collection Error", err.Error())
		return
	}

	data.Id = types.StringValue(collection.Id)
	data.Metadata, _ = types.MapValueFrom(ctx, types.StringType, collection.Metadata)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CollectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CollectionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	collection, err := r.getCollectionByID(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Get Collection Error", err.Error())
		return
	}

	data.Id = types.StringValue(collection.Id)
	data.Metadata, _ = types.MapValueFrom(ctx, types.StringType, collection.Metadata)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CollectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CollectionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state CollectionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	metadata := make(map[string]interface{}, len(data.Metadata.Elements()))
	for k, v := range data.Metadata.Elements() {
		strVal, ok := v.(types.String)
		if !ok || strVal.IsNull() {
			continue
		}
		metadata[k] = strVal.ValueString()
	}

	collection, err := r.updateCollection(ctx, state.Id.ValueString(), hautechapi.CollectionsControllerUpdateMetadataV1JSONRequestBody{Overwrite: metadata})
	if err != nil {
		resp.Diagnostics.AddError("Update Collection Error", err.Error())
		return
	}

	data.Id = types.StringValue(collection.Id)
	data.Metadata, _ = types.MapValueFrom(ctx, types.StringType, collection.Metadata)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CollectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r *CollectionResource) getCollectionByID(ctx context.Context, id string) (*hautechapi.CollectionEntity, error) {
	resp, err := r.client.CollectionsControllerGetCollectionV1WithResponse(ctx, id)
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

func (r *CollectionResource) createCollection(ctx context.Context, params hautechapi.CreateCollectionParamsDto) (*hautechapi.CollectionEntity, error) {
	resp, err := r.client.CollectionsControllerCreateCollectionV1WithResponse(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated || resp.JSON201 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return resp.JSON201, nil
}

func (r *CollectionResource) updateCollection(ctx context.Context, id string, params hautechapi.CollectionsControllerUpdateMetadataV1JSONRequestBody) (*hautechapi.CollectionEntity, error) {
	resp, err := r.client.CollectionsControllerUpdateMetadataV1WithResponse(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("collection not found")
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return r.getCollectionByID(ctx, id)
}
