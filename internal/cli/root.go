package cli

import (
	"context"

	"github.com/gregmulvaney/omni/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type app struct {
	v        *viper.Viper
	cfgFlags *pflag.FlagSet
}

func Execute(ctx context.Context) error {
	return rootCmd().ExecuteContext(ctx)
}

func rootCmd() *cobra.Command {
	a := &app{
		v:        viper.New(),
		cfgFlags: config.RegisterFlags(),
	}
	root := &cobra.Command{
		Use:          "omni",
		Short:        "Root command for the app",
		SilenceUsage: true,
	}

	root.PersistentFlags().AddFlagSet(a.cfgFlags)

	root.AddCommand(configCmd(a))

	return root
}
