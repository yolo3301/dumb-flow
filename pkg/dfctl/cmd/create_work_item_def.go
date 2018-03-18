package cmd

import (
	"log"
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
		for _, v := range configs {
			d := strings.Index(v, "=")
			if d == -1 || d == len(v)-1 {
				log.Panicf("Invalid config '%v'", v)
			}

			configs[v[:d]] = v[d+1:]
		}

		log.Print("To be implemented...")
	},
}

func init() {
	CreateCommand.AddCommand(CreateWorkItemDefCmd)

	CreateWorkItemDefCmd.Flags().StringVar(&createWorkItemDefWorkflowName, "workflow", "", "The workflow name")
	CreateWorkItemDefCmd.MarkFlagRequired("workflow")
	CreateWorkItemDefCmd.Flags().StringVar(&createWorkItemDefWorkItemName, "name", "", "The work item name")
	CreateWorkItemDefCmd.MarkFlagRequired("name")
	createWorkflowDefWorkflowConfigs = *CreateWorkItemDefCmd.Flags().StringSlice("configs", nil, "The work item configs")
}
