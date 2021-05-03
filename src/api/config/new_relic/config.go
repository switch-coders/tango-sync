package new_relic

import (
	"os"
	"strings"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func Init() (*newrelic.Application, error) {
	scope := os.Getenv("SCOPE")
	if strings.HasSuffix(scope, "master") {
		return setupMasterEnvironment()
	}

	if strings.HasSuffix(scope, "beta") {
		return setupBetaEnvironment()
	}

	return newrelic.NewApplication(
		newrelic.ConfigAppName("master.tango-sync"),
		newrelic.ConfigLicense("42f8705fa3ae7feb35621fc56b973691ee4cNRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigEnabled(false),
	)
}

func setupMasterEnvironment() (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName("master.tango-sync"),
		newrelic.ConfigLicense("42f8705fa3ae7feb35621fc56b973691ee4cNRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(config *newrelic.Config) {
			config.ErrorCollector.IgnoreStatusCodes = []int{400, 401, 403, 404, 405, 423, 424, 429}
		},
	)
}

func setupBetaEnvironment() (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName("beta.tango-sync"),
		newrelic.ConfigLicense("b3e5f010b616aa389d1f34a771498824f710NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(config *newrelic.Config) {
			config.ErrorCollector.IgnoreStatusCodes = []int{400, 401, 403, 404, 405, 423, 424, 429}
		},
	)
}