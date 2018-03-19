package cmd

import (
	"github.com/spf13/cobra"
)

var CreateCommand = &cobra.Command{
	Use:   "create",
	Short: "create stuff",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(CreateCommand)
}
