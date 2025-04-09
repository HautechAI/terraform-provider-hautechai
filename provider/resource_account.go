package provider

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	hautechapi "hautech/api"
)

var ErrAccountNotFound = errors.New("account not found")

type AccountResource struct {
	client *hautechapi.ClientWithResponses
}

type GroupAssignment struct {
	GroupID types.String `tfsdk:"group_id"`
	Role    types.String `tfsdk:"role"`
}

type AccountResourceModel struct {
	Id    types.String     `tfsdk:"id"`
	Alias types.String     `tfsdk:"alias"`
	Group *GroupAssignment `tfsdk:"group"`
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
			"alias": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"group": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"group_id": schema.StringAttribute{Required: true},
					"role":     schema.StringAttribute{Required: true},
				},
			},
		},
	}
}

func (r *AccountResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData != nil {
		r.client = req.ProviderData.(*Context).Client
	}
}

func (r *AccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AccountResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	acc, err := r.getOrCreateAccount(ctx, data.Alias.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Get or create Account Error", err.Error())
		return
	}

	data.Id = types.StringValue(acc.Id)

	if data.Group != nil {
		err := r.assignGroup(ctx, acc.Id, data.Group.GroupID.ValueString(), hautechapi.AddAccountToGroupControllerParamsDtoRole(data.Group.Role.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Group Assignment Error", err.Error())
			return
		}
	}

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
	var data, state AccountResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	acc, err := r.getOrCreateAccount(ctx, data.Alias.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Get or create Account Error", err.Error())
		return
	}

	// Remove old group if it was set and now is nil or changed
	if state.Group != nil && (data.Group == nil || data.Group.GroupID.ValueString() != state.Group.GroupID.ValueString()) {
		err := r.removeGroup(ctx, acc.Id, state.Group.GroupID.ValueString(), hautechapi.RemoveAccountFromGroupControllerParamsDtoRole(state.Group.Role.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Remove Group Error", err.Error())
			return
		}
	}

	// Assign new group if defined
	if data.Group != nil {
		err := r.assignGroup(ctx, acc.Id, data.Group.GroupID.ValueString(), hautechapi.AddAccountToGroupControllerParamsDtoRole(data.Group.Role.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Assign Group Error", err.Error())
			return
		}
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
		return nil, ErrAccountNotFound
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

func (r *AccountResource) getOrCreateAccount(ctx context.Context, alias string) (*hautechapi.AccountEntity, error) {
	if alias != "" {
		acc, err := r.getAccountByAlias(ctx, alias)
		if err != nil && !errors.Is(err, ErrAccountNotFound) {
			return nil, err
		}
		if acc != nil {
			return acc, nil
		}
	}
	return r.createAccount(ctx, alias)
}

func (r *AccountResource) assignGroup(ctx context.Context, accountID, groupID string, role hautechapi.AddAccountToGroupControllerParamsDtoRole) error {
	body := hautechapi.AddAccountToGroupControllerParamsDto{
		AccountId: accountID,
		Role:      role,
	}
	resp, err := r.client.GroupsControllerAddAccountV1WithResponse(ctx, groupID, body)
	if err != nil || resp.StatusCode() >= 300 {
		return fmt.Errorf("failed to assign group %s for account %s (%s): %v - %s", groupID, accountID, role, err, string(resp.Body))
	}
	return nil
}

func (r *AccountResource) removeGroup(ctx context.Context, accountID, groupID string, role hautechapi.RemoveAccountFromGroupControllerParamsDtoRole) error {
	body := hautechapi.RemoveAccountFromGroupControllerParamsDto{
		AccountId: accountID,
		Role:      role,
	}
	resp, err := r.client.GroupsControllerRemoveAccountV1WithResponse(ctx, groupID, body)
	if err != nil || resp.StatusCode() >= 300 {
		return fmt.Errorf("failed to remove group %s for account %s: %v - %s", groupID, accountID, err, string(resp.Body))
	}
	return nil
}
