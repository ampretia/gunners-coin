package main

import (
	"fmt"

	"github.com/ampretia/gunners-coin/internal/config"
	"github.com/ampretia/gunners-coin/internal/logger"
	views "github.com/ampretia/gunners-coin/views"
	fscnode "github.com/hyperledger-labs/fabric-smart-client/node"
	fabric "github.com/hyperledger-labs/fabric-smart-client/platform/fabric/sdk"

	viewregistry "github.com/hyperledger-labs/fabric-smart-client/platform/view"

	sdk "github.com/hyperledger-labs/fabric-token-sdk/token/sdk"
)

func main() {

	fmt.Println("Hello world!")
	fmt.Printf("Got an env variable: %q\n", config.Foo)
	logger.Infof("I can even use a logger package!")

	n := fscnode.New()
	n.InstallSDK(fabric.NewSDK(n))
	n.InstallSDK(sdk.NewSDK(n))

	n.Execute(func() error {
		registry := viewregistry.GetRegistry(n)
		if err := registry.RegisterFactory("issue", &views.IssueCashViewFactory{}); err != nil {
			return err
		}
		if err := registry.RegisterFactory("issued", &views.ListIssuedTokensViewFactory{}); err != nil {
			return err
		}

		return nil
	})
}
