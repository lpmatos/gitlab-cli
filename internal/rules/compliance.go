package rules

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// ComplianceRule - struct to represent a compliance rule.
type ComplianceRule struct {
	Name        string
	Description string
}

// Local function - check values of compliance rule struct.
func (rule *ComplianceRule) validateComplienceRule() {
	if rule.Name == "" {
		log.Warn("Compliance rule name is empty")
	}
	if rule.Description == "" {
		log.Warn("Compliance rule description is empty")
	}
}

// Local function - parse compliance rule struct and return a format string.
func (rule *ComplianceRule) parseComplienceRule() string {
	rule.validateComplienceRule()
	var info = ` 
Rule name: %s
	
Rule description: 
%s`
	return fmt.Sprintf(info, rule.Name, rule.Description)
}

// Show function - print the compliance rule struct formatted.
func (rule *ComplianceRule) Show() {
	fmt.Println(rule.parseComplienceRule())
}
