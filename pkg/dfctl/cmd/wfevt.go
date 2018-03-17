package cmd

import (
	"github.com/spf13/cobra"
)

func NewEventCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "event",
		Short: "event",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	return command
}
