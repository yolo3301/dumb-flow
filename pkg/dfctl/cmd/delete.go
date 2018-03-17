package cmd

import (
	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "delete",
		Short: "delete stuff",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(NewWorkflowDefCmd())
	command.AddCommand(NewWorkItemDefCmd())
	command.AddCommand(NewWorkflowExecCmd())
	command.AddCommand(NewWorkItemExecCmd())
	command.AddCommand(NewEventCmd())

	return command
}
