package block

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lpmatos/glabby/internal/commands/compliance/options"
	"github.com/lpmatos/glabby/internal/config"
	"github.com/lpmatos/glabby/internal/helpers"
	types "github.com/lpmatos/glabby/internal/json"
	log "github.com/sirupsen/logrus"
)

// Local function - check block sonar.
func validateBlockSonar(block bool) {
	if block || strings.ToLower(config.GetEnv("BLOCK_SONAR", "")) == "true" {
		log.Fatal("Sonar job was blocked in pipeline")
		os.Exit(1)
	}
	log.Fatal("Sonar job allows failure")
}

// Local function - rule to check sonar request GET status and decide if we block or not.
func ruleSonarBlock(status string, block bool) {
	if status == "ERROR" {
		helpers.ASCIIMessage("-----------------", "banner", "red")
		log.Warn("This project isn't in the compliance")
		validateBlockSonar(block)
	} else {
		helpers.ASCIIMessage("-----------------", "banner", "green")
		fmt.Println("This project is in the compliance")
	}
}

// SonarAnalyzer function - sonar analyzer complience... we analyze the sonar endpoint (/api/qualitygates/project_status).
func SonarAnalyzer(sonar *options.SonarOptions) {
	log.Info("Running sonar compliance analyzer")

	result := fmt.Sprintf("%s/api/qualitygates/project_status?projectKey=%s", sonar.URL, sonar.Project)

	log.Printf("Sonar URL: %s", result)

	var (
		client = &http.Client{
			Timeout: time.Second * 15,
		}
		requestJSON types.SonarRequestJSON
	)

	log.Info("Doing GET request")

	req, error := http.NewRequest(http.MethodGet, result, nil)

	if error != nil {
		log.Fatal(error)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(sonar.Token, "")

	resp, error := client.Do(req)

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if error != nil {
		log.Fatal(error)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Bad status code...")
	}

	log.Info("Success requests. Read body json content")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal([]byte(body), &requestJSON); err != nil {
		log.Fatal(err)
	}

	data := types.JSON{Data: &requestJSON}
	data.PrettyPrintJSON(sonar.Pretty, body)

	fmt.Printf("\n\n")

	log.Info("Analyzing sonar json metrics...")

	ruleSonarBlock(requestJSON.ProjectStatus.Status, sonar.Block)
}
