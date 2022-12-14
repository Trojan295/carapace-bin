package cmd

import (
	"github.com/spf13/cobra"
)

var p4Cmd = &cobra.Command{
	Use:     "p4",
	Short:   "Import from and submit to Perforce repositories",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {

	rootCmd.AddCommand(p4Cmd)
}
