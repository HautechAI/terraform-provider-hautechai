package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"net/url"

	hautechapi "hautech/api"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &hautechProvider{}

type ProviderConfig struct {
	ApiUrl string
}

type hautechProvider struct {
	apiUrl string
}

func New(cfg ProviderConfig) func() provider.Provider {
	return func() provider.Provider {
		return &hautechProvider{
			apiUrl: cfg.ApiUrl,
		}
	}
}

type providerModel struct {
	ApiToken types.String `tfsdk:"api_token"`
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
		},
	}
}

func createClient(apiURL string, apiToken string) (*hautechapi.APIClient, error) {
	parsedUrl, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	cfg := hautechapi.NewConfiguration()
	cfg.Host = parsedUrl.Host
	cfg.Scheme = parsedUrl.Scheme
	cfg.DefaultHeader["Authorization"] = "Bearer " + apiToken

	return hautechapi.NewAPIClient(cfg), nil
}

type ProviderContext struct {
	Client *hautechapi.APIClient
}

func (p *hautechProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config providerModel

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := createClient(p.apiUrl, config.ApiToken.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error while creating http client", err.Error())
		return
	}

	resp.ResourceData = &ProviderContext{Client: client}
	resp.DataSourceData = &ProviderContext{Client: client}
}

func (p *hautechProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

func (p *hautechProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAvailablePermissionsDataSource,
	}
}
