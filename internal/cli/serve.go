package cli

import (
	"github.com/gregmulvaney/omni/internal/app"
	"github.com/spf13/cobra"
)

func serveCmd(a *appConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Run Omni",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := a.loadConfig()
			if err != nil {
				return err
			}

			app, err := app.Bootstrap(cfg)
			app.Log.Info("Omni starting...")

			<-cmd.Context().Done()

			app.Log.Info("Shutting down")

			return nil
		},
	}
}
