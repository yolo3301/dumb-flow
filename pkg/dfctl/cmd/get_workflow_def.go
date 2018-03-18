package cmd

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
	dfclient "github.com/yolo3301/dumb-flow/pkg/df-client"
)

var getWorkflowDefWorkflowName string

var GetWorkflowDefCmd = &cobra.Command{
	Use:     "workflow-def",
	Aliases: []string{"wfdef"},
	Short:   "get workflow def",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := dfclient.NewDumbflowClient()
		if err != nil {
			log.Fatal(err.Error())
		}

		def, err := client.GetWorkflowDef(getWorkflowDefWorkflowName)
		if err != nil {
			log.Panic(err.Error())
		}

		content, err := json.Marshal(def)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Print(string(content))
	},
}

func init() {
	GetCmd.AddCommand(GetWorkflowDefCmd)

	GetWorkflowDefCmd.Flags().StringVar(&getWorkflowDefWorkflowName, "name", "", "The workflow name")
	GetWorkflowDefCmd.MarkFlagRequired("name")
}
