package main

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"hautech/provider"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/hautech/api",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(provider.ProviderConfig{
		ApiUrl: getApiUrl(),
	}), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
