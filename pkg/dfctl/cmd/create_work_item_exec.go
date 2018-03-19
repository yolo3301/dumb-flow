package cmd

import (
	"github.com/spf13/cobra"
)

var createWorkItemExecWorkflowName string
var createWorkItemExecWorkflowExecID string
var createWorkItemExecWorkItemName string

var CreateWorkItemExecCmd = &cobra.Command{
	Use:     "workitem-exec",
	Aliases: []string{"wiexec"},
	Short:   "create work item exec",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func init() {
	CreateCommand.AddCommand(CreateWorkItemExecCmd)
}
