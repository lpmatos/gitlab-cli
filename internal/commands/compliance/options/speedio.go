package options

import (
	log "github.com/sirupsen/logrus"

	"github.com/lpmatos/glabby/internal/helpers"
	"github.com/spf13/cobra"
)

// SpeedioOptions struct - represents all options that we need in speedio compliance.
type SpeedioOptions struct {
	File   string
	Pretty bool
	Block  bool
}

// ValidateSpeedioOptions function - check speedio options to compliance.
func (se *SpeedioOptions) ValidateSpeedioOptions() {
	if se.File == "" {
		log.Fatal("Speedio file parameter is empty")
	}
	if !helpers.FileExists(se.File) {
		log.Fatal("Speedio file does not exist")
	}
}

// AddSpeedioOptions function - add speedio options to a top level command (compliance).
func (se *SpeedioOptions) AddSpeedioOptions(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&se.File, "file", "f", "sitespeed-results/data/performance.json",
		"this will define speedio json file used in compliance")
	cmd.Flags().BoolVarP(&se.Pretty, "pretty", "p", false,
		"enable pretty json print")
	cmd.Flags().BoolVarP(&se.Pretty, "block", "b", false,
		"enable block to speedio job in pipeline")
}
