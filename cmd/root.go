package cmd

import (
	"github.com/lpmatos/glabby/internal/commands"
	"github.com/lpmatos/glabby/internal/commands/compliance"
	"github.com/lpmatos/glabby/internal/logging"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var silence bool

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "glabby",
	Short: "GitLab CLI (Compliance)",
	Long: `Description:

Glabby is a CLI library to interact with your pipeline and make/analyze some stuffs.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	logging.Setup()

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&silence, "silence", "s", false, "enable silence mod without logs in stdout terminal.")

	compliance.AddComplianceCommands(rootCmd)
	commands.AddVersionCommand(rootCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	silence, _ := rootCmd.Flags().GetBool("silence")
	if silence {
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
