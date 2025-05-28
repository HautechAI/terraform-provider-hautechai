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

type StackResource struct {
	client *hautechapi.ClientWithResponses
}

type StackResourceModel struct {
	ID       types.String `tfsdk:"id"`
	Metadata types.Map    `tfsdk:"metadata"`
	Items    types.List   `tfsdk:"items"`
}

func NewStackResource() resource.Resource {
	return &StackResource{}
}

func (r *StackResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_stack"
}

func (r *StackResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage Hautech Stacks.",
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

func (r *StackResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData != nil {
		r.client = req.ProviderData.(*Context).Client
	}
}

func (r *StackResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StackResourceModel
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

	stack, err := r.createStack(ctx, hautechapi.CreateStackParamsDto{Metadata: &metadata})
	if err != nil {
		resp.Diagnostics.AddError("Create Stack Error", err.Error())
		return
	}

	data.ID = types.StringValue(stack.Id)

	if data.Metadata.IsNull() && len(stack.Metadata) == 0 {
		data.Metadata = types.MapNull(types.StringType)
	} else {
		data.Metadata, _ = types.MapValueFrom(ctx, types.StringType, stack.Metadata)
	}

	if !data.Items.IsNull() {
		var itemIds []string
		data.Items.ElementsAs(ctx, &itemIds, false)
		err := r.updateStackItems(ctx, stack.Id, nil, itemIds)
		if err != nil {
			resp.Diagnostics.AddError("Add items Error", err.Error())
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StackResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StackResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	stack, err := r.getStackByID(ctx, data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Get Stack Error", err.Error())
		return
	}

	data.ID = types.StringValue(stack.Id)

	if data.Metadata.IsNull() && len(stack.Metadata) == 0 {
		data.Metadata = types.MapNull(types.StringType)
	} else {
		data.Metadata, _ = types.MapValueFrom(ctx, types.StringType, stack.Metadata)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StackResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan StackResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state StackResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
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

	stack, err := r.updateStack(ctx, state.ID.ValueString(), hautechapi.StacksControllerUpdateMetadataV1JSONRequestBody{Overwrite: metadata})
	if err != nil {
		resp.Diagnostics.AddError("Update Stack Error", err.Error())
		return
	}

	plan.ID = types.StringValue(stack.Id)

	if plan.Metadata.IsNull() && len(stack.Metadata) == 0 {
		plan.Metadata = types.MapNull(types.StringType)
	} else {
		plan.Metadata, _ = types.MapValueFrom(ctx, types.StringType, stack.Metadata)
	}

	var newItems, oldItems []string
	plan.Items.ElementsAs(ctx, &newItems, false)
	state.Items.ElementsAs(ctx, &oldItems, false)

	toAdd := diffStrings(newItems, oldItems)
	toRemove := diffStrings(oldItems, newItems)

	if len(toAdd) > 0 || len(toRemove) > 0 {
		err := r.updateStackItems(ctx, stack.Id, toRemove, toAdd)
		if err != nil {
			resp.Diagnostics.AddError("Update items Error", err.Error())
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *StackResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r *StackResource) getStackByID(ctx context.Context, id string) (*hautechapi.StackEntity, error) {
	resp, err := r.client.StacksControllerGetStackV1WithResponse(ctx, id)
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

func (r *StackResource) createStack(ctx context.Context, params hautechapi.CreateStackParamsDto) (*hautechapi.StackEntity, error) {
	resp, err := r.client.StacksControllerCreateStackV1WithResponse(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}
	if resp.StatusCode() != http.StatusCreated || resp.JSON201 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}
	return resp.JSON201, nil
}

func (r *StackResource) updateStack(ctx context.Context, id string, params hautechapi.StacksControllerUpdateMetadataV1JSONRequestBody) (*hautechapi.StackEntity, error) {
	resp, err := r.client.StacksControllerUpdateMetadataV1WithResponse(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}
	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("stack not found")
	}
	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}
	return r.getStackByID(ctx, id)
}

func (r *StackResource) updateStackItems(ctx context.Context, StackID string, removeIDs, addIDs []string) error {
	if len(removeIDs) > 0 {
		removeBody := hautechapi.StacksControllerRemoveItemsV1JSONRequestBody{ItemIds: removeIDs}
		resp, err := r.client.StacksControllerRemoveItemsV1WithResponse(ctx, StackID, removeBody)
		if err != nil || resp.StatusCode() >= 300 {
			return fmt.Errorf("failed to remove items: %v - %s", err, string(resp.Body))
		}
	}
	if len(addIDs) > 0 {
		addBody := hautechapi.StacksControllerAddItemsV1JSONRequestBody{ItemIds: addIDs}
		resp, err := r.client.StacksControllerAddItemsV1WithResponse(ctx, StackID, addBody)
		if err != nil || resp.StatusCode() >= 300 {
			return fmt.Errorf("failed to add items: %v - %s", err, string(resp.Body))
		}
	}
	return nil
}
