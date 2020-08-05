package compliance

import (
	"github.com/lpmatos/glabby/internal/commands/compliance/block"
	"github.com/lpmatos/glabby/internal/commands/compliance/options"
	"github.com/lpmatos/glabby/internal/helpers"
	"github.com/lpmatos/glabby/internal/rules"
	"github.com/spf13/cobra"
)

// Rule speedio description
var ruleSpeedio = &rules.ComplianceRule{
	Name: "Compliance to block speedio in pipeline",
	Description: `
* After speedio execution, it generates a json file.
* This json file was some performance metrics.
* We analyzed this file looking for the total score value.
* If this value is less than 85, then the project is considered out of compliance.
* If 'BLOCK_SPEEDIO' is enabled, then this job will fail and the pipeline fail.
`,
}

// Function to build speedio command, add our options and attach to a top level command (compliance).
func addSpeedio(compliance *cobra.Command) {
	se := options.SpeedioOptions{}
	speedio := &cobra.Command{
		Use:   "speedio",
		Short: "Run speedio compliance",
		Long: `Description:

This command run speedio compliance to a pipeline in GitLab.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			ruleSpeedio.Show()
			helpers.ASCIIMessage("-----------------", "banner", "green")
			se.ValidateSpeedioOptions()
			block.SpeedioAnalyzer(&se)
		},
	}
	se.AddSpeedioOptions(speedio)
	compliance.AddCommand(speedio)
}
