package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	hautechapi "hautech/api"
)

var _ resource.Resource = &AccountResource{}

type AccountResource struct {
	client *hautechapi.ClientWithResponses
}

type AccountResourceModel struct {
	Id    types.String `tfsdk:"id"`
	Alias types.String `tfsdk:"alias"`
}

func NewAccountResource() resource.Resource {
	return &AccountResource{}
}

func (r *AccountResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account"
}

func (r *AccountResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"alias": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *AccountResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*ProviderContext).Client
}

func (r *AccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AccountResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...) // get alias from input
	if resp.Diagnostics.HasError() {
		return
	}

	id, diags := r.getOrCreateAccount(ctx, data.Alias.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = types.StringValue(id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...) // save id to state
	tflog.Trace(ctx, "Account resource created or found", map[string]any{"id": data.Id.ValueString(), "alias": data.Alias.ValueString()})
}

func (r *AccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AccountResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...) // read current state
	if resp.Diagnostics.HasError() {
		return
	}

	getResp, err := r.client.AccountsControllerGetAccountV1WithResponse(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("API Error - Read Account", fmt.Sprintf("failed to read account: %v", err))
		return
	}

	if getResp.StatusCode() == http.StatusOK && getResp.JSON200 != nil {
		data.Id = types.StringValue(getResp.JSON200.Id)
	} else if getResp.StatusCode() == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	} else {
		resp.Diagnostics.AddError("API Error - Read Account", fmt.Sprintf("unexpected status code: %d", getResp.StatusCode()))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...) // update state
	tflog.Trace(ctx, "Account resource read", map[string]any{"id": data.Id.ValueString(), "alias": data.Alias.ValueString()})
}

func (r *AccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AccountResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...) // get new alias
	if resp.Diagnostics.HasError() {
		return
	}

	id, diags := r.getOrCreateAccount(ctx, data.Alias.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = types.StringValue(id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...) // update state
	tflog.Trace(ctx, "Account resource updated or created", map[string]any{"id": data.Id.ValueString(), "alias": data.Alias.ValueString()})
}

func (r *AccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
	tflog.Trace(ctx, "Account resource deleted")
}

func (r *AccountResource) getOrCreateAccount(ctx context.Context, alias string) (string, diag.Diagnostics) {
	var diags diag.Diagnostics

	getResp, err := r.client.AccountsControllerGetAccountByAliasV1WithResponse(ctx, alias)
	if err != nil {
		diags.AddError("API Error - Get Account by Alias", fmt.Sprintf("failed to get account by alias: %v", err))
		return "", diags
	}

	if getResp.StatusCode() == http.StatusOK && getResp.JSON200 != nil {
		return getResp.JSON200.Id, diags
	}

	if getResp.StatusCode() == http.StatusNotFound {
		createParams := hautechapi.CreateAccountParamsDto{Alias: &alias}
		createResp, err := r.client.AccountsControllerCreateAccountV1WithResponse(ctx, createParams)
		if err != nil {
			diags.AddError("API Error - Create Account", fmt.Sprintf("failed to create account: %v", err))
			return "", diags
		}
		if createResp.StatusCode() != http.StatusCreated || createResp.JSON201 == nil {
			diags.AddError("API Error - Create Account", "unexpected response from account creation")
			return "", diags
		}
		return createResp.JSON201.Id, diags
	}

	diags.AddError("API Error - Get Account by Alias", fmt.Sprintf("unexpected status code: %d", getResp.StatusCode()))
	return "", diags
}
