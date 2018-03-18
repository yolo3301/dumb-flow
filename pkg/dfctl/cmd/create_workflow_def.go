package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	dfclient "github.com/yolo3301/dumb-flow/pkg/df-client"
)

var createWorkflowDefWorkflowName string
var createWorkflowDefWorkflowConfigs []string

var CreateWorkflowDefCmd = &cobra.Command{
	Use:     "workflow-def",
	Aliases: []string{"wfdef"},
	Short:   "create workflow def",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := dfclient.NewDumbflowClient()
		if err != nil {
			log.Fatal(err.Error())
		}

		configs := make(map[string]string)
		for _, v := range createWorkflowDefWorkflowConfigs {
			d := strings.Index(v, "=")
			if d == -1 || d == len(v)-1 {
				log.Panicf("Invalid config '%v'", v)
			}

			configs[v[:d]] = v[d+1:]
		}

		err = client.CreateWorkflowDef(createWorkflowDefWorkflowName, configs)
		if err != nil {
			log.Panic(err.Error())
		}

		fmt.Printf("Created workflow '%v'", createWorkflowDefWorkflowName)
	},
}

func init() {
	CreateCommand.AddCommand(CreateWorkflowDefCmd)

	CreateWorkflowDefCmd.Flags().StringVar(&createWorkflowDefWorkflowName, "workflow-name", "", "The workflow name")
	CreateWorkflowDefCmd.MarkFlagRequired("workflow-name")
	createWorkflowDefWorkflowConfigs = *CreateWorkflowDefCmd.Flags().StringSlice("configs", nil, "The workflow configs")
}
