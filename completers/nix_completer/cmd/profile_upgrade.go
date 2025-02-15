package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var profile_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade packages using their most recent flake",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_upgradeCmd).Standalone()

	profile_upgradeCmd.Flags().String("profile", "", "The profile to operate on")
	profileCmd.AddCommand(profile_upgradeCmd)

	addEvaluationFlags(profile_upgradeCmd)
	addFlakeFlags(profile_upgradeCmd)
	addLoggingFlags(profile_upgradeCmd)

	carapace.Gen(profile_upgradeCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionDirectories(),
	})
}
