package cmd

import (
	"github.com/spf13/cobra"
)

func NewWorkItemDefCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "work item def",
		Short: "work item def",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	return command
}
