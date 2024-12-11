package cmd

import (
	"github.com/diki-haryadi/gateway-gopher/internal"
	"github.com/spf13/cobra"
)

var (
	gatewayCmd = &cobra.Command{
		Use:              "gw",
		Short:            "Gateway",
		Long:             "Gateway",
		PersistentPreRun: gatewayPreRun,
		RunE:             runGateway,
	}
)

var (
	Gateway = internal.NewGateway()
)

func GatewayCmd() *cobra.Command {
	return gatewayCmd
}

func gatewayPreRun(cmd *cobra.Command, args []string) {}

func runGateway(cmd *cobra.Command, args []string) error {
	internal.App(Gateway)
	return nil
}
