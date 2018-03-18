package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	dfclient "github.com/yolo3301/dumb-flow/pkg/df-client"
)

var deleteWorkflowExecWorkflowName string
var deleteWorkflowExecWorkflowExecID string

var DeleteWorkflowExecCmd = &cobra.Command{
	Use:     "workflow-exec",
	Aliases: []string{"wfexec"},
	Short:   "delete workflow exec",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := dfclient.NewDumbflowClient()
		if err != nil {
			log.Fatal(err.Error())
		}

		err = client.DeleteWorkflowExec(deleteWorkflowDefWorkflowName, deleteWorkflowExecWorkflowExecID)
		if err != nil {
			log.Panic(err.Error())
		}

		fmt.Printf("Deleted workflow exec '%v'", deleteWorkflowExecWorkflowExecID)
	},
}

func init() {
	DeleteCmd.AddCommand(DeleteWorkflowExecCmd)

	DeleteWorkflowExecCmd.Flags().StringVar(&deleteWorkflowDefWorkflowName, "workflow-name", "", "The workflow name")
	DeleteWorkflowExecCmd.MarkFlagRequired("workflow-name")

	DeleteWorkflowExecCmd.Flags().StringVar(&deleteWorkflowExecWorkflowExecID, "workflow-exec", "", "The workflow exec id")
	DeleteWorkflowExecCmd.MarkFlagRequired("workflow-exec")

	// todo add cascade option
}
