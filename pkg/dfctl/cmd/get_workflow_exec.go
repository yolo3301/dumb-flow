package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	dfclient "github.com/yolo3301/dumb-flow/pkg/df-client"
)

var getWorkflowExecWorkflowName string
var getWorkflowExecWorkflowExecID string

var GetWorkflowExecCmd = &cobra.Command{
	Use:     "workflow-exec",
	Aliases: []string{"wfexec"},
	Short:   "get workflow exec",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := dfclient.NewDumbflowClient()
		if err != nil {
			log.Fatal(err.Error())
		}

		exec, err := client.GetWorkflowExec(getWorkflowDefWorkflowName, getWorkflowExecWorkflowExecID)
		if err != nil {
			log.Panic(err.Error())
		}

		content, err := json.Marshal(exec)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Print(string(content))
	},
}

func init() {
	GetCmd.AddCommand(GetWorkflowExecCmd)

	GetWorkflowExecCmd.Flags().StringVar(&getWorkflowDefWorkflowName, "workflow-name", "", "The workflow name")
	GetWorkflowExecCmd.MarkFlagRequired("workflow-name")

	GetWorkflowExecCmd.Flags().StringVar(&getWorkflowExecWorkflowExecID, "workflow-exec", "", "The workflow exec id")
	GetWorkflowExecCmd.MarkFlagRequired("workflow-exec")

	// todo add expand option
}
