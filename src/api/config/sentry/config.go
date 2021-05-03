package sentry

import (
	"log"
	"os"
	"strings"

	"github.com/getsentry/sentry-go"
)

func Init() {
	scope := os.Getenv("SCOPE")
	if strings.HasSuffix(scope, "master") {
		setupMasterEnvironment()
		return
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://eb7dc2cbc3fe4ce9b9a008b22a24b55f@o610821.ingest.sentry.io/5748014",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

func setupMasterEnvironment() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://af5eaa774f6f4476ac5aea4991889796@o610885.ingest.sentry.io/5748020",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
