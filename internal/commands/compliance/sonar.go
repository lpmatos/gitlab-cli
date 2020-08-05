package compliance

import (
	"github.com/lpmatos/glabby/internal/commands/compliance/block"
	"github.com/lpmatos/glabby/internal/commands/compliance/options"
	"github.com/lpmatos/glabby/internal/helpers"
	"github.com/lpmatos/glabby/internal/rules"
	"github.com/spf13/cobra"
)

// Rule sonar description.
var ruleSonar = &rules.ComplianceRule{
	Name: "Compliance to block sonar in pipeline",
	Description: `
* After the sonar execution we GET the endpoint /api/qualitygates/project_status.
* This endpoint returns quality gates in sonar for that project.
* If the response return a json 'status' with the 'ERROR' value, then the project is considered out of compliance.
* If 'BLOCK_SONAR' is enabled, then the job will fail and the pipeline fail.
`,
}

// Function to build sonar command, add our options and attach to a top level command (compliance).
func addSonar(topLevel *cobra.Command) {
	so := options.SonarOptions{}
	sonar := &cobra.Command{
		Use:   "sonar",
		Short: "Run sonar compliance",
		Long: `Description:

This command run sonar compliance to a pipeline in GitLab.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			ruleSonar.Show()
			helpers.ASCIIMessage("-----------------", "banner", "green")
			so.ValidateSonarOptions()
			block.SonarAnalyzer(&so)
		},
	}
	so.AddSonarOptions(sonar)
	topLevel.AddCommand(sonar)
}
