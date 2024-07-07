package cmd

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	tmcfg "github.com/tendermint/tendermint/config"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
)

// Execute executes the root command of an application. It handles creating a
// server context object with the appropriate server and client objects injected
// into the underlying stdlib Context. It also handles adding core CLI flags,
// specifically the logging flags. It returns an error upon execution failure.
func Execute(rootCmd *cobra.Command, envPrefix string, defaultHome string) error {
	// Create and set a client.Context on the command's Context. During the pre-run
	// of the root command, a default initialized client.Context is provided to
	// seed child command execution with values such as AccountRetriver, Keyring,
	// and a Tendermint RPC. This requires the use of a pointer reference when
	// getting and setting the client.Context. Ideally, we utilize
	// https://github.com/spf13/cobra/pull/1118.
	srvCtx := server.NewDefaultContext()
	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &client.Context{})
	ctx = context.WithValue(ctx, server.ServerContextKey, srvCtx)

	rootCmd.PersistentFlags().String(flags.FlagLogLevel, zerolog.InfoLevel.String(), "The logging level (trace|debug|info|warn|error|fatal|panic)")
	rootCmd.PersistentFlags().String(flags.FlagLogFormat, tmcfg.LogFormatPlain, "The logging format (json|plain)")
	rootCmd.PersistentFlags().Bool(flags.FlagLogNoColor, true, "Disable colored logs")
	rootCmd.PersistentFlags().Int(flags.FlagLogMaxSize, 1024, "Maximum space occupied by a single log file")
	rootCmd.PersistentFlags().Int(flags.FlagLogMaxBackups, 168, "The maximum number of log files to be retained for split storage")
	rootCmd.PersistentFlags().Int(flags.FlagLogMaxAge, 7, "The maximum retention time of the split and stored log files, in days")

	executor := tmcli.PrepareBaseCmd(rootCmd, envPrefix, defaultHome)
	return executor.ExecuteContext(ctx)
}
