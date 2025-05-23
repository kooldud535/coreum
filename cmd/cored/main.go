package main

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/CoreumFoundation/coreum/v6/app"
	"github.com/CoreumFoundation/coreum/v6/cmd/cored/cosmoscmd"
)

const coreumEnvPrefix = "CORED"

func main() {
	network, err := cosmoscmd.PreProcessFlags()
	if err != nil {
		fmt.Printf("Error processing chain id flag, err: %s", err)
		os.Exit(1)
	}
	network.SetSDKConfig()
	app.ChosenNetwork = network

	rootCmd := cosmoscmd.NewRootCmd()
	cosmoscmd.OverwriteDefaultChainIDFlags(rootCmd)
	rootCmd.PersistentFlags().String(flags.FlagChainID, string(app.DefaultChainID), "The network chain ID")
	if err := svrcmd.Execute(rootCmd, coreumEnvPrefix, app.DefaultNodeHome); err != nil {
		//nolint:errcheck // we are already exiting the app so we don't check error.
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
