package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	dfclient "github.com/yolo3301/dumb-flow/pkg/df-client"
)

var createWorkflowExecWorkflowName string

var CreateWorkflowExecCmd = &cobra.Command{
	Use:     "workflow-exec",
	Aliases: []string{"wfexec"},
	Short:   "create workflow exec",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := dfclient.NewDumbflowClient()
		if err != nil {
			log.Fatal(err.Error())
		}

		id, err := client.CreateWorkflowExec(createWorkflowExecWorkflowName)
		if err != nil {
			log.Panic(err.Error())
		}

		fmt.Printf("Created workflow '%v' exec id = '%v'\n", createWorkflowExecWorkflowName, id)
	},
}

func init() {
	CreateCommand.AddCommand(CreateWorkflowExecCmd)

	CreateWorkflowExecCmd.Flags().StringVar(&createWorkflowExecWorkflowName, "workflow-name", "", "The workflow name")
	CreateWorkflowExecCmd.MarkFlagRequired("workflow-name")
}
