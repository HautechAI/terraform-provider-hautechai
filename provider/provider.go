package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	hautechapi "hautech/api"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Config struct {
	ApiUrl string
}

type Context struct {
	Client *hautechapi.ClientWithResponses
}

type hautechProvider struct {
	apiUrl string
}

func New(cfg Config) func() provider.Provider {
	return func() provider.Provider {
		return &hautechProvider{
			apiUrl: cfg.ApiUrl,
		}
	}
}

type providerModel struct {
	ApiToken types.String `tfsdk:"api_token"`
	ApiUrl   types.String `tfsdk:"api_url"`
}

func (p *hautechProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "hautech"
}

func (p *hautechProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Hautech provider for calling the Hautech API.",
		Attributes: map[string]schema.Attribute{
			"api_token": schema.StringAttribute{
				Optional:    true,
				Description: "API token. Read how to get it: https://docs.hautech.ai/getting-started",
			},
			"api_url": schema.StringAttribute{
				Optional:    true,
				Description: "API URL. If not provided, the default URL will be used.",
			},
		},
	}
}

func (p *hautechProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config providerModel

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Use user-provided API URL if available, otherwise use the default
	apiUrl := p.apiUrl
	if !config.ApiUrl.IsNull() && config.ApiUrl.ValueString() != "" {
		apiUrl = config.ApiUrl.ValueString()
	}

	client, err := createClient(apiUrl, config.ApiToken.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error while creating http client", err.Error())
		return
	}

	contextClient := &Context{Client: client}
	resp.ResourceData = contextClient
	resp.DataSourceData = contextClient
}

func (p *hautechProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAccountResource,
		NewCollectionResource,
		NewGroupResource,
		NewStackResource,
	}
}

func (p *hautechProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAvailablePermissionsDataSource,
		NewAccountBalanceDataSource,
		NewAccountDataSource,
		NewCollectionDataSource,
		NewStackDataSource,
	}
}

func createClient(apiURL string, apiToken string) (*hautechapi.ClientWithResponses, error) {
	authEditor := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+apiToken)
		return nil
	}

	client, err := hautechapi.NewClientWithResponses(apiURL, hautechapi.WithRequestEditorFn(authEditor))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return client, nil
}
