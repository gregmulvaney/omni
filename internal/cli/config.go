package cli

import (
	"fmt"

	"github.com/gregmulvaney/omni/internal/config"
	"github.com/spf13/cobra"
)

func configCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage omni configuration",
	}

	cmd.AddCommand(configGenerateCmd())

	return cmd
}

func configGenerateCmd() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a config file with the default values",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var path string

			if len(args) == 1 {
				path = args[0]
			} else {
				path = config.DefaultPath
			}

			if err := config.WriteDefault(path, force); err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "wrote config file %s\n", path)

			return nil
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing config file")

	return cmd
}
