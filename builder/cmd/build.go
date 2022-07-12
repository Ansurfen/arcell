package cmd

import (
	"arcell/build"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		flags := make(map[string]any)
		if len(args) >= 1 {
			build.Build(args, flags)
		}
		println("Args require more than or equal 1")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
