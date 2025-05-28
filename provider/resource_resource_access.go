package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	hautechapi "hautech/api"
)

type ResourceAccessResource struct {
	client *hautechapi.ClientWithResponses
}

type ResourceAccessResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ResourceId    types.String `tfsdk:"resource_id"`
	PrincipalId   types.String `tfsdk:"principal_id"`
	PrincipalType types.String `tfsdk:"principal_type"`
	Access        types.String `tfsdk:"access"`
}

func NewResourceAccessResource() resource.Resource {
	return &ResourceAccessResource{}
}

func (r *ResourceAccessResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_access"
}

func (r *ResourceAccessResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Grants access to a resource for a principal.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the resource access grant.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"resource_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the resource to grant access to.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"principal_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the principal (account, user, group) to grant access to.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"principal_type": schema.StringAttribute{
				Required:    true,
				Description: "The type of principal (account, user, group).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"access": schema.StringAttribute{
				Required:    true,
				Description: "The access level to grant.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *ResourceAccessResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData != nil {
		r.client = req.ProviderData.(*Context).Client
	}
}

func (r *ResourceAccessResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ResourceAccessResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	access := data.Access.ValueString()
	if access != "maintainer" && access != "writer" && access != "reader" {
		resp.Diagnostics.AddError(
			"Invalid access value",
			fmt.Sprintf("The access value must be one of: maintainer, writer, reader. Got: %s", access),
		)
		return
	}

	err := r.grantAccess(ctx, data.ResourceId.ValueString(), data.PrincipalId.ValueString(),
		hautechapi.GrantAccessControllerParamsPrincipalType(data.PrincipalType.ValueString()),
		data.Access.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error granting access", err.Error())
		return
	}

	data.Id = types.StringValue(fmt.Sprintf("%s:%s:%s:%s",
		data.ResourceId.ValueString(),
		data.PrincipalType.ValueString(),
		data.PrincipalId.ValueString(),
		data.Access.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResourceAccessResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ResourceAccessResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// We don't have a direct way to check if the access is still granted
	// So we'll just keep the state as is
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResourceAccessResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ResourceAccessResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	access := data.Access.ValueString()
	if access != "maintainer" && access != "writer" && access != "reader" {
		resp.Diagnostics.AddError(
			"Invalid access value",
			fmt.Sprintf("The access value must be one of: maintainer, writer, reader. Got: %s", access),
		)
		return
	}

	// First revoke the existing access
	err := r.revokeAccess(ctx, data.ResourceId.ValueString(), data.PrincipalId.ValueString(),
		hautechapi.RevokeAccessControllerParamsDtoPrincipalType(data.PrincipalType.ValueString()), data.Access.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error revoking access during update", err.Error())
		return
	}

	// Then grant the new access
	err = r.grantAccess(ctx, data.ResourceId.ValueString(), data.PrincipalId.ValueString(),
		hautechapi.GrantAccessControllerParamsPrincipalType(data.PrincipalType.ValueString()),
		data.Access.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error granting access during update", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResourceAccessResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ResourceAccessResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.revokeAccess(ctx, data.ResourceId.ValueString(), data.PrincipalId.ValueString(),
		hautechapi.RevokeAccessControllerParamsDtoPrincipalType(data.PrincipalType.ValueString()), data.Access.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error revoking access", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *ResourceAccessResource) grantAccess(ctx context.Context, resourceId, principalId string, principalType hautechapi.GrantAccessControllerParamsPrincipalType, access string) error {
	body := hautechapi.GrantAccessControllerParams{
		Access:        access,
		PrincipalId:   principalId,
		PrincipalType: principalType,
	}
	resp, err := r.client.AccessControllerGrantAccessV1WithResponse(ctx, resourceId, body)
	if err != nil {
		return fmt.Errorf("failed to call API: %w", err)
	}
	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}
	return nil
}

func (r *ResourceAccessResource) revokeAccess(ctx context.Context, resourceId, principalId string, principalType hautechapi.RevokeAccessControllerParamsDtoPrincipalType, access string) error {
	body := hautechapi.RevokeAccessControllerParamsDto{
		Access:        access,
		PrincipalId:   principalId,
		PrincipalType: principalType,
	}
	resp, err := r.client.AccessControllerRevokeAccessV1WithResponse(ctx, resourceId, body)
	if err != nil {
		return fmt.Errorf("failed to call API: %w", err)
	}
	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}
	return nil
}
