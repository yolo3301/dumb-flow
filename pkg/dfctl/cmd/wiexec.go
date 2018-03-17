package cmd

import (
	"github.com/spf13/cobra"
)

func NewWorkItemExecCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "work item exec",
		Short: "work item exec",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	return command
}
