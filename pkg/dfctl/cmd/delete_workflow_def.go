package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	dfclient "github.com/yolo3301/dumb-flow/pkg/df-client"
)

var deleteWorkflowDefWorkflowName string

var DeleteWorkflowDefCmd = &cobra.Command{
	Use:     "workflow-def",
	Aliases: []string{"wfdef"},
	Short:   "delete workflow def",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := dfclient.NewDumbflowClient()
		if err != nil {
			log.Fatal(err.Error())
		}

		err = client.DeleteWorkflowDef(deleteWorkflowDefWorkflowName)
		if err != nil {
			log.Panic(err.Error())
		}

		fmt.Printf("Deleted workflow '%v'\n", deleteWorkflowDefWorkflowName)
	},
}

func init() {
	DeleteCmd.AddCommand(DeleteWorkflowDefCmd)

	DeleteWorkflowDefCmd.Flags().StringVar(&deleteWorkflowDefWorkflowName, "workflow-name", "", "The workflow name")
	DeleteWorkflowDefCmd.MarkFlagRequired("workflow-name")
}
