package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "touch",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "change only the access time")
	rootCmd.Flags().StringP("date", "d", "", "parse STRING and use it instead of current time")
	rootCmd.Flags().BoolP("f", "f", false, "(ignored)")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("m", "m", false, "change only the modification time")
	rootCmd.Flags().BoolP("no-create", "c", false, "do not create any files")
	rootCmd.Flags().BoolP("no-dereference", "h", false, "affect each symbolic link instead of any referenced file")
	rootCmd.Flags().StringP("reference", "r", "", "use this file's times instead of current time")
	rootCmd.Flags().StringP("t", "t", "", "use [[CC]YY]MMDDhhmm[.ss] instead of current time")
	rootCmd.Flags().String("time", "", "change the specified time")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"reference": carapace.ActionFiles(""),
		"time":      carapace.ActionValues("access", "atime", "modify", "mtime", "use"),
	})
}
