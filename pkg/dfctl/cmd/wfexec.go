package cmd

import (
	"github.com/spf13/cobra"
)

func NewWorkflowExecCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "workflow exec",
		Short: "workflow exec",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	return command
}
