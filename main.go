package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/switch-coders/tango-sync/src/api/app"
	"github.com/switch-coders/tango-sync/src/api/config"
	"github.com/switch-coders/tango-sync/src/api/config/new_relic"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
	"github.com/switch-coders/tango-sync/src/api/infrastructure/dependencies"
	"github.com/switch-coders/tango-sync/src/api/infrastructure/jobs"
)

func main() {
	port := os.Getenv("PORT")
	config.SetupEnvironment()

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	handlers := dependencies.Start()
	router := gin.Default()
	nrApp, _ := new_relic.Init()
	router.Use(nrgin.Middleware(nrApp))
	router.Use(infrastructure.SetNRTransaction)
	app.ConfigureMappings(router, handlers)
	jobs.Schedule()
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
