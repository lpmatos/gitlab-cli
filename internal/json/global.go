package json

import (
	"encoding/json"
	"fmt"

	"github.com/lpmatos/glabby/internal/helpers"
	log "github.com/sirupsen/logrus"
)

// JSON struct that abstract a data mapping a interface.
type JSON struct {
	Data interface{}
}

// Local function - pretty print a JSON.
func (data *JSON) prettyPrint() {
	pretty, err := json.MarshalIndent(data.Data, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n\n%s", pretty)
}

// PrettyPrintJSON function - check pretty print option and show the information.
func (data *JSON) PrettyPrintJSON(pretty bool, value []byte) {
	helpers.ASCIIMessage("JSON", "", "")
	if pretty {
		data.prettyPrint()
	} else {
		fmt.Printf("\n\n%s", value)
	}
}
