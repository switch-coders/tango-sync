package config

import (
	"os"
	"strings"

	"github.com/switch-coders/tango-sync/src/api/config/sentry"
)

const (
	HerokuURL         = "https://tango-sync.herokuapp.com"
	HerokuBetaURL     = "https://tango-sync-beta.herokuapp.com"
	LocalURL          = "http://localhost:8080"
	TangoBaseURL      = "https://tiendas.axoft.com/api/Aperture"
	TiendaNubeBaseURL = "https://api.tiendanube.com/v1/%s"
)

func SetupEnvironment() {
	_ = os.Setenv("BASE_URL", LocalURL)
	_ = os.Setenv("TANGO_BASE_URL", TangoBaseURL)
	_ = os.Setenv("TN_BASE_URL", TiendaNubeBaseURL)

	scope := os.Getenv("SCOPE")
	if strings.HasSuffix(scope, "master") {
		setupMasterEnvironment()
	}

	if strings.HasSuffix(scope, "beta") {
		setupBetaEnvironment()
	}

	_ = os.Setenv("TN_SECRET", "MHgj0k7L7w37KtDprnPvHniE6LNcmMh9RbGebTf2OIfSg8Rv")
	_ = os.Setenv("TN_APP_ID", "2736")
}

func setupMasterEnvironment() {
	_ = os.Setenv("BASE_URL", HerokuURL)
	sentry.Init()
}

func setupBetaEnvironment() {
	_ = os.Setenv("BASE_URL", HerokuBetaURL)
	sentry.Init()
}
