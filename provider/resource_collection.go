package provider

import (
	"context"
	"fmt"
	"net/http"

	hautechapi "hautech/api"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CollectionResource struct {
	client *hautechapi.ClientWithResponses
}

type CollectionResourceModel struct {
	ID       types.String `tfsdk:"id"`
	Metadata types.Map    `tfsdk:"metadata"`
	Items    types.List   `tfsdk:"items"`
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
			"items": schema.ListAttribute{
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

	data.ID = types.StringValue(collection.Id)
	data.Metadata, _ = types.MapValueFrom(ctx, types.StringType, collection.Metadata)

	if !data.Items.IsNull() {
		var itemIds []string
		data.Items.ElementsAs(ctx, &itemIds, false)
		err := r.updateCollectionItems(ctx, collection.Id, nil, itemIds)
		if err != nil {
			resp.Diagnostics.AddError("Add items Error", err.Error())
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CollectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CollectionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	collection, err := r.getCollectionByID(ctx, data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Get Collection Error", err.Error())
		return
	}

	data.ID = types.StringValue(collection.Id)
	data.Metadata, _ = types.MapValueFrom(ctx, types.StringType, collection.Metadata)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CollectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CollectionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...) // desired
	if resp.Diagnostics.HasError() {
		return
	}

	var state CollectionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...) // current
	if resp.Diagnostics.HasError() {
		return
	}

	metadata := make(map[string]interface{}, len(plan.Metadata.Elements()))
	for k, v := range plan.Metadata.Elements() {
		strVal, ok := v.(types.String)
		if !ok || strVal.IsNull() {
			continue
		}
		metadata[k] = strVal.ValueString()
	}

	collection, err := r.updateCollection(ctx, state.ID.ValueString(), hautechapi.CollectionsControllerUpdateMetadataV1JSONRequestBody{Overwrite: metadata})
	if err != nil {
		resp.Diagnostics.AddError("Update Collection Error", err.Error())
		return
	}

	plan.ID = types.StringValue(collection.Id)
	plan.Metadata, _ = types.MapValueFrom(ctx, types.StringType, collection.Metadata)

	var newItems, oldItems []string
	plan.Items.ElementsAs(ctx, &newItems, false)
	state.Items.ElementsAs(ctx, &oldItems, false)

	toAdd := diffStrings(newItems, oldItems)
	toRemove := diffStrings(oldItems, newItems)

	if len(toAdd) > 0 || len(toRemove) > 0 {
		err := r.updateCollectionItems(ctx, collection.Id, toRemove, toAdd)
		if err != nil {
			resp.Diagnostics.AddError("Update items Error", err.Error())
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
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

func (r *CollectionResource) updateCollectionItems(ctx context.Context, collectionID string, removeIDs, addIDs []string) error {
	if len(removeIDs) > 0 {
		removeBody := hautechapi.CollectionsControllerRemoveItemsV1JSONRequestBody{ItemIds: removeIDs}
		resp, err := r.client.CollectionsControllerRemoveItemsV1WithResponse(ctx, collectionID, removeBody)
		if err != nil || resp.StatusCode() >= 300 {
			return fmt.Errorf("failed to remove items: %v - %s", err, string(resp.Body))
		}
	}
	if len(addIDs) > 0 {
		addBody := hautechapi.CollectionsControllerAddItemsV1JSONRequestBody{ItemIds: addIDs}
		resp, err := r.client.CollectionsControllerAddItemsV1WithResponse(ctx, collectionID, addBody)
		if err != nil || resp.StatusCode() >= 300 {
			return fmt.Errorf("failed to add items: %v - %s", err, string(resp.Body))
		}
	}
	return nil
}
