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
	_ = os.Setenv("TN_SECRET", "MHgj0k7L7w37KtDprnPvHniE6LNcmMh9RbGebTf2OIfSg8Rv")
	_ = os.Setenv("TN_APP_ID", "2736")

	scope := os.Getenv("SCOPE")
	if strings.HasSuffix(scope, "master") {
		setupMasterEnvironment()
	}

	if strings.HasSuffix(scope, "beta") {
		setupBetaEnvironment()
	}

}

func setupMasterEnvironment() {
	_ = os.Setenv("BASE_URL", HerokuURL)
	_ = os.Setenv("TN_SECRET", "gOXthoAwKJRaEdT7rE3hd6J3y0WljELi0osgu60JoggaL419")
	_ = os.Setenv("TN_APP_ID", "2953")
	sentry.Init()
}

func setupBetaEnvironment() {
	_ = os.Setenv("BASE_URL", HerokuBetaURL)
	_ = os.Setenv("TN_SECRET", "6rpM1u2BgX54kAUGmUI8OW4yDKT1pPOFS2L4i5vFozI3mIaB")
	_ = os.Setenv("TN_APP_ID", "3189")
	sentry.Init()
}
