package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Current const
const version = "1.3.0"

// AddVersionCommand function - top level command.
func AddVersionCommand(topLevel *cobra.Command) {
	version := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Version outputs the version of CLI",
		Long:    `Version outputs the version of the glabby binary that is in use.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(currentVersion())
		},
	}
	topLevel.AddCommand(version)
}

// Local function - parse current version and return a formatted string.
func currentVersion() string {
	return fmt.Sprintf("glabby version %s", version)
}
