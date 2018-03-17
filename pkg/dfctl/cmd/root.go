package cmd

import (
	"github.com/spf13/cobra"
)

const (
	// CLIName is the name of the CLI
	CLIName = "dfctl"
)

func NewCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   CLIName,
		Short: "dfctl is the command line for dumbflow",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(NewCreateCmd())
	command.AddCommand(NewDeleteCmd())
	command.AddCommand(NewGetCmd())

	return command
}
