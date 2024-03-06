package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Short:  "gin frame",
	Long:   "gin frame project",
	Run:    func(cmd *cobra.Command, args []string) {
		start()
	},
}