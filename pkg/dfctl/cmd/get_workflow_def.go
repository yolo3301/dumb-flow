package cmd

import (
	"encoding/json"
	"fmt"
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

		var content []byte

		if getWorkflowDefWorkflowName != "" {
			def, err := client.GetWorkflowDef(getWorkflowDefWorkflowName)
			if err != nil {
				log.Panic(err.Error())
			}

			content, err = json.Marshal(def)
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			defs, err := client.GetWorkflowDefs()
			if err != nil {
				log.Panic(err.Error())
			}

			content, err = json.Marshal(defs)
			if err != nil {
				log.Fatal(err.Error())
			}
		}

		fmt.Println(string(content))
	},
}

func init() {
	GetCmd.AddCommand(GetWorkflowDefCmd)

	GetWorkflowDefCmd.Flags().StringVar(&getWorkflowDefWorkflowName, "workflow-name", "", "The workflow name")
}
