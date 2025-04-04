package main

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"hautech/config"
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
		Address: config.GetProviderAddress(),
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(provider.Config{
		ApiUrl: config.GetApiUrl(),
	}), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
