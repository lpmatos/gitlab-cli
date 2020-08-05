package options

import (
	log "github.com/sirupsen/logrus"

	"github.com/lpmatos/glabby/internal/config"
	"github.com/lpmatos/glabby/internal/helpers"
	"github.com/spf13/cobra"
)

// SonarOptions struct - represents all options that we need in sonar compliance.
type SonarOptions struct {
	URL     string
	Token   string
	Project string
	Pretty  bool
	Block   bool
}

// ValidateSonarOptions function - check sonar options to compliance.
func (so *SonarOptions) ValidateSonarOptions() {
	if helpers.Empty(so.URL) {
		log.Fatal("Sonar url is empty")
	}
	if helpers.Empty(so.Token) {
		log.Fatal("Sonar token is empty")
	}
	if helpers.Empty(so.Project) {
		log.Fatal("Sonar project is empty")
	}
	if !helpers.IsURL(so.URL) {
		log.Fatal("Sonar invalid url")
	}
}

// AddSonarOptions function - add sonar options to a top level command (compliance).
func (so *SonarOptions) AddSonarOptions(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&so.URL, "url", "u",
		config.GetEnv("SONAR_URL", ""), "this will define the Sonar URL")
	cmd.Flags().StringVarP(&so.Token, "token", "t",
		config.GetEnv("SONAR_TOKEN", ""), "this will define the Sonar Token")
	cmd.Flags().StringVarP(&so.Project, "project", "P",
		config.GetEnv("SONAR_PROJECT_KEY", ""), "this will define the Sonar Project Key")
	cmd.Flags().BoolVarP(&so.Pretty, "pretty", "p", false,
		"enable pretty json print")
	cmd.Flags().BoolVarP(&so.Block, "block", "b", false,
		"enable block to sonar job in pipeline")
}
