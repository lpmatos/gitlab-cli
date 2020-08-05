package block

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/lpmatos/glabby/internal/commands/compliance/options"
	"github.com/lpmatos/glabby/internal/config"
	"github.com/lpmatos/glabby/internal/helpers"
	types "github.com/lpmatos/glabby/internal/json"
	log "github.com/sirupsen/logrus"
)

// Local function - check block speedio.
func validateBlockSpeedio(block bool) {
	if block || strings.ToLower(config.GetEnv("BLOCK_SPEEDIO", "")) == "true" {
		log.Fatal("Speedio job was blocked in pipeline")
		os.Exit(1)
	}
	log.Fatal("Speedio job allows failure")
}

// Local function - rule to check score of speedio performance json file and decide if we block or not.
func ruleSpeedioBlock(score int, block bool) {
	if score < 85 {
		helpers.ASCIIMessage("-----------------", "banner", "red")
		log.Warn("This project isn't in the compliance")
		validateBlockSpeedio(block)
	} else {
		helpers.ASCIIMessage("-----------------", "banner", "green")
		fmt.Println("This project is in the compliance")
	}
}

// SpeedioAnalyzer function - speedio analyzer complience... we analyze the content of performance json file created in speedio job.
func SpeedioAnalyzer(speedio *options.SpeedioOptions) {
	log.Info("Running speedio compliance analyzer")

	var speedioJSON []types.SpeedioJSON

	log.Printf("Reading performance json file: %s", speedio.File)

	file, err := os.Open(speedio.File)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.Info("Read performance json content")

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(content), &speedioJSON)

	data := types.JSON{Data: &speedioJSON}
	data.PrettyPrintJSON(speedio.Pretty, content)

	score := speedioJSON[0].Metrics[0].Value

	fmt.Print("\n\n")

	message := "Score: " + strconv.Itoa(score)

	helpers.ASCIIMessage(message, "", "")

	fmt.Printf("\n\n")

	log.Printf("Analyzing speedio json score...")

	ruleSpeedioBlock(score, speedio.Block)
}
