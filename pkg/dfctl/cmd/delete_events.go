package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	dfclient "github.com/yolo3301/dumb-flow/pkg/df-client"
)

var deleteEventsWorkflowName string
var deleteEventsWorkflowExec string

var DeleteEventsCmd = &cobra.Command{
	Use:   "event",
	Short: "delete events",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := dfclient.NewDumbflowClient()
		if err != nil {
			log.Fatal(err.Error())
		}

		err = client.DeleteEvents(deleteEventsWorkflowName, deleteEventsWorkflowExec)
		if err != nil {
			log.Panic(err.Error())
		}

		fmt.Printf("Deleted events for workflow '%v' exec id = '%v'\n", deleteEventsWorkflowName, deleteEventsWorkflowExec)
	},
}

func init() {
	DeleteCmd.AddCommand(DeleteEventsCmd)

	DeleteEventsCmd.Flags().StringVar(&deleteEventsWorkflowName, "workflow-name", "", "The workflow name")
	DeleteEventsCmd.MarkFlagRequired("workflow-name")

	DeleteEventsCmd.Flags().StringVar(&deleteEventsWorkflowExec, "workflow-exec", "", "The workflow exec id")
	DeleteEventsCmd.MarkFlagRequired("workflow-exec")
}
