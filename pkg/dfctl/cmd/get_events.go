package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	dfclient "github.com/yolo3301/dumb-flow/pkg/df-client"
)

var getEventsWorkflowName string
var getEventsWorkflowExecID string

var GetEventsCmd = &cobra.Command{
	Use:   "event",
	Short: "get events",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := dfclient.NewDumbflowClient()
		if err != nil {
			log.Fatal(err.Error())
		}

		events, err := client.GetEvents(getEventsWorkflowName, getEventsWorkflowExecID)
		if err != nil {
			log.Panic(err.Error())
		}

		content, err := json.Marshal(events)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(string(content))
	},
}

func init() {
	GetCmd.AddCommand(GetEventsCmd)

	GetEventsCmd.Flags().StringVar(&getEventsWorkflowName, "workflow-name", "", "The workflow name")
	GetEventsCmd.MarkFlagRequired("workflow-name")

	GetEventsCmd.Flags().StringVar(&getEventsWorkflowExecID, "workflow-exec", "", "The workflow exec id")
	GetEventsCmd.MarkFlagRequired("workflow-exec")
}
