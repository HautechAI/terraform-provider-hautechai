package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	hautechapi "hautech/api"
)

type AccountResource struct {
	client *hautechapi.ClientWithResponses
}

type AccountResourceModel struct {
	Id        types.String `tfsdk:"id"`
	AccountId types.String `tfsdk:"account_id"`
	Alias     types.String `tfsdk:"alias"`
}

func NewAccountResource() resource.Resource {
	return &AccountResource{}
}

func (r *AccountResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account"
}

func (r *AccountResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage Hautech accounts.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"alias": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *AccountResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData != nil {
		r.client = req.ProviderData.(*ProviderContext).Client
	}
}

func (r *AccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AccountResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	acc, err := r.getOrCreateAccount(ctx, data.Alias.ValueString(), data.AccountId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Get or create Account Error", err.Error())
		return
	}

	data.Id = types.StringValue(acc.Id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AccountResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	acc, err := r.getAccountByID(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Get Account Error", err.Error())
		return
	}

	data.Id = types.StringValue(acc.Id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AccountResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	acc, err := r.getOrCreateAccount(ctx, data.Alias.ValueString(), data.AccountId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Get or create Account Error", err.Error())
		return
	}

	data.Id = types.StringValue(acc.Id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r *AccountResource) getAccountByID(ctx context.Context, id string) (*hautechapi.AccountEntity, error) {
	resp, err := r.client.AccountsControllerGetAccountV1WithResponse(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("account not found")
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return resp.JSON200, nil
}

func (r *AccountResource) getAccountByAlias(ctx context.Context, alias string) (*hautechapi.AccountEntity, error) {
	resp, err := r.client.AccountsControllerGetAccountByAliasV1WithResponse(ctx, alias)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("account not found")
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return resp.JSON200, nil
}

func (r *AccountResource) createAccount(ctx context.Context, alias string) (*hautechapi.AccountEntity, error) {
	params := hautechapi.CreateAccountParamsDto{}
	if alias != "" {
		params.Alias = &alias
	}

	resp, err := r.client.AccountsControllerCreateAccountV1WithResponse(ctx, params)

	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated || resp.JSON201 == nil {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return resp.JSON201, nil
}

func (r *AccountResource) getOrCreateAccount(ctx context.Context, alias, id string) (*hautechapi.AccountEntity, error) {
	if id != "" {
		acc, err := r.getAccountByID(ctx, id)

		return acc, err
	}

	if alias != "" {
		acc, err := r.getAccountByAlias(ctx, alias)
		if err != nil && err.Error() != "account not found" {
			return nil, err
		}

		if acc != nil {
			return acc, nil
		}
	}

	acc, err := r.createAccount(ctx, alias)

	return acc, err
}
