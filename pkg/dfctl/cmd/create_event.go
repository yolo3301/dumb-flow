package cmd

import (
	"fmt"
	"log"

	"github.com/yolo3301/dumb-flow/pkg/df-model"

	"github.com/spf13/cobra"
	dfclient "github.com/yolo3301/dumb-flow/pkg/df-client"
)

var createEventWorkflowName string
var createEventWorkflowExecID string
var createEventWorkItemName string
var createEventWorkItemExecID string
var createEventPayload string
var createEventType string

var CreateEventCmd = &cobra.Command{
	Use:   "event",
	Short: "create event",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := dfclient.NewDumbflowClient()
		if err != nil {
			log.Fatal(err.Error())
		}

		id, err := client.CreateEvent(
			createEventWorkflowName,
			createEventWorkflowExecID,
			createEventWorkItemName,
			createEventWorkItemExecID,
			createEventPayload,
			model.EventType(createEventType))

		if err != nil {
			log.Panic(err.Error())
		}

		fmt.Printf("Created event id = '%v'", id)
	},
}

func init() {
	CreateCommand.AddCommand(CreateEventCmd)

	CreateEventCmd.Flags().StringVar(&createEventWorkflowName, "workflow-name", "", "The workflow name")
	CreateEventCmd.MarkFlagRequired("workflow-name")

	CreateEventCmd.Flags().StringVar(&createEventWorkflowExecID, "workflow-exec", "", "The workflow exec id")
	CreateEventCmd.MarkFlagRequired("workflow-exec")

	CreateEventCmd.Flags().StringVar(&createEventWorkItemName, "workitem-name", "", "The work item name")

	CreateEventCmd.Flags().StringVar(&createEventWorkItemExecID, "workitem-exec", "", "The work item exec id")

	CreateEventCmd.Flags().StringVar(&createEventType, "type", "", "The event type")
	CreateEventCmd.MarkFlagRequired("type")

	CreateEventCmd.Flags().StringVar(&createEventPayload, "payload", "", "The event payload")
}
