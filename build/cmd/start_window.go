package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Short:  "gin frame",
	Long:   "gin frame project",
	Run:    func(cmd *cobra.Command, args []string) {
		start()
	},
}