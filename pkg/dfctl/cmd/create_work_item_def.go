package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var createWorkItemDefWorkflowName string
var createWorkItemDefWorkItemName string
var createWorkItemDefWorkItemConfigs []string

var CreateWorkItemDefCmd = &cobra.Command{
	Use:     "workitem-def",
	Aliases: []string{"widef"},
	Short:   "create work item def",
	Run: func(cmd *cobra.Command, args []string) {
		configs := make(map[string]string)
		for _, v := range createWorkflowDefWorkflowConfigs {
			d := strings.Index(v, "=")
			if d == -1 || d == len(v)-1 {
				fmt.Printf("Invalid config '%v'", v)
			}

			configs[v[:d]] = v[d+1:]
		}

		fmt.Print("To be implemented...")
	},
}

func init() {
	CreateCommand.AddCommand(CreateWorkItemDefCmd)

	CreateWorkItemDefCmd.Flags().StringVar(&createWorkItemDefWorkflowName, "workflow-name", "", "The workflow name")
	CreateWorkItemDefCmd.MarkFlagRequired("workflow-name")
	CreateWorkItemDefCmd.Flags().StringVar(&createWorkItemDefWorkItemName, "workitem-name", "", "The work item name")
	CreateWorkItemDefCmd.MarkFlagRequired("workitem-name")
	createWorkflowDefWorkflowConfigs = *(CreateWorkItemDefCmd.Flags().StringSlice("configs", nil, "The work item configs"))
}
