package cli

import (
	"context"

	"github.com/gregmulvaney/omni/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type appConfig struct {
	v        *viper.Viper
	cfgFlags *pflag.FlagSet
}

func Execute(ctx context.Context) error {
	return rootCmd().ExecuteContext(ctx)
}

func rootCmd() *cobra.Command {
	a := &appConfig{
		v:        viper.New(),
		cfgFlags: config.RegisterFlags(),
	}

	root := &cobra.Command{
		Use:          "omni",
		Short:        "Root command for the app",
		SilenceUsage: true,
	}

	root.PersistentFlags().AddFlagSet(a.cfgFlags)

	root.AddCommand(
		configCmd(a),
		serveCmd(a),
	)

	return root
}

func (a *appConfig) loadConfig() (*config.Config, error) {
	return config.Load(a.v, a.cfgFlags, "")
}
