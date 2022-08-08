package main

import (
	"fmt"

	fscnode "github.com/hyperledger-labs/fabric-smart-client/node"
	fabric "github.com/hyperledger-labs/fabric-smart-client/platform/fabric/sdk"
	sdk "github.com/hyperledger-labs/fabric-token-sdk/token/sdk"

	"github.com/ampretia/gunners-coin/internal/config"
	"github.com/ampretia/gunners-coin/internal/logger"
	views "github.com/ampretia/gunners-coin/views"
	viewregistry "github.com/hyperledger-labs/fabric-smart-client/platform/view"
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
		if err := registry.RegisterFactory("transfer", &views.TransferViewFactory{}); err != nil {
			return err
		}
		if err := registry.RegisterFactory("redeem", &views.RedeemViewFactory{}); err != nil {
			return err
		}
		if err := registry.RegisterFactory("swap", &views.SwapInitiatorViewFactory{}); err != nil {
			return err
		}
		if err := registry.RegisterFactory("unspent", &views.ListUnspentTokensViewFactory{}); err != nil {
			return err
		}
		registry.RegisterResponder(&views.AcceptCashView{}, &views.IssueCashView{})
		registry.RegisterResponder(&views.AcceptCashView{}, &views.TransferView{})

		return nil
	})
}
