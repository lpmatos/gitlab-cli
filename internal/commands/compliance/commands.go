package compliance

import "github.com/spf13/cobra"

// AddComplianceCommands function - top level command.
func AddComplianceCommands(topLevel *cobra.Command) {
	compliance := &cobra.Command{
		Use:     "compliance",
		Aliases: []string{"c"},
		Short:   "Run compliance in pipeline",
		Long: `Description:
	
Compliance to your GitLab pipeline.
				`,
	}
	addSonar(compliance)
	addSpeedio(compliance)
	topLevel.AddCommand(compliance)
}
