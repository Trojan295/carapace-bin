package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_approversCmd = &cobra.Command{
	Use:   "approvers [<id> | <branch>] [flags]",
	Short: "List eligible approvers for merge requests in any state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_approversCmd).Standalone()

	mrCmd.AddCommand(mr_approversCmd)

	carapace.Gen(mr_approversCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_approversCmd, ""),
	)
}
