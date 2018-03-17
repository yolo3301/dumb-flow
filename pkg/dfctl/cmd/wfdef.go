package cmd

import (
	"github.com/spf13/cobra"
)

func NewWorkflowDefCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "workflow def",
		Short: "workflow def",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	return command
}
