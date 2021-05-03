package config

import (
	"os"
	"strings"

	"github.com/switch-coders/tango-sync/src/api/config/sentry"
)

const (
	HerokuURL         = "https://tango-sync.herokuapp.com"
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
}

func setupMasterEnvironment() {
	_ = os.Setenv("BASE_URL", HerokuURL)
	sentry.Init()
}
