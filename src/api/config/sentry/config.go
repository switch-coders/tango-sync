package sentry

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func Init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://53d9e3685ab34e748243c98a75c133f1:cea305334e3840b1b680b0d0ad66cac9@o570194.ingest.sentry.io/5716869",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
