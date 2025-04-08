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

type GroupResource struct {
	client *hautechapi.ClientWithResponses
}

type GroupResourceModel struct {
	ID types.String `tfsdk:"id"`
}

func NewGroupResource() resource.Resource {
	return &GroupResource{}
}

func (r *GroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (r *GroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage Hautech Groups.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *GroupResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData != nil {
		r.client = req.ProviderData.(*Context).Client
	}
}

func (r *GroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	group, err := r.createGroup(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Create Group Error", err.Error())
		return
	}

	data.ID = types.StringValue(group.Id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	group, err := r.getGroupByID(ctx, data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Get Group Error", err.Error())
		return
	}

	data.ID = types.StringValue(group.Id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Nothing for now
}

func (r *GroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r *GroupResource) getGroupByID(ctx context.Context, id string) (*hautechapi.GroupEntity, error) {
	resp, err := r.client.GroupsControllerGetGroupV1WithResponse(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("group not found")
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return resp.JSON200, nil
}

func (r *GroupResource) createGroup(ctx context.Context) (*hautechapi.GroupEntity, error) {
	resp, err := r.client.GroupsControllerCreateGroupV1WithResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated || resp.JSON201 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return resp.JSON201, nil
}
