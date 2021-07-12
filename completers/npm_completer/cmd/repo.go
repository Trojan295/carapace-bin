package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Open package repository page in the browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()
	repoCmd.Flags().String("browser", "", "browser to user")
	repoCmd.Flags().StringP("workspace", "w", "", "Enable running a command in the context of the given workspace")
	repoCmd.Flags().Bool("workspaces", false, "Enable running a command in the context fo all workspaces")

	rootCmd.AddCommand(repoCmd)

	carapace.Gen(repoCmd).FlagCompletion(carapace.ActionMap{
		"browser": carapace.ActionFiles(),
	})

	carapace.Gen(repoCmd).PositionalAnyCompletion(
		action.ActionPackageNames(repoCmd),
	)
}
