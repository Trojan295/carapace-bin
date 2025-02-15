package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sshKey_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of currently authenticated user's SSH keys.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_listCmd).Standalone()

	sshKey_listCmd.Flags().StringP("page", "p", "", "Page number")
	sshKey_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page")
	sshKey_listCmd.Flags().Bool("show-id", false, "Shows IDs of SSH Keys")
	sshKeyCmd.AddCommand(sshKey_listCmd)
}
