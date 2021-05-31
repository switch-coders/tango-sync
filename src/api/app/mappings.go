package app

import (
	"github.com/gin-gonic/gin"
	"github.com/switch-coders/tango-sync/src/api/infrastructure/dependencies"
	"net/http"
)

func ConfigureMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	configureApiMappings(router, handlers)
	configureJobsMappings(router, handlers)
	configureConsumersMappings(router, handlers)
}

func configureApiMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	router.LoadHTMLGlob("./templates/*")
	router.Static("/assets", "./assets")

	group := router.Group("/")
	group.GET("ping", handlers.Get.Handle)
	group.GET("index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	group.GET("register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})
	group.POST("registration", handlers.Registration.Handle)
	group.POST("integration", handlers.Integration.Handle)
}

func configureJobsMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	groupJobs := router.Group("/jobs/sync")
	groupJobs.POST("stock", handlers.SyncStock.Handle)
	groupJobs.POST("price", handlers.SyncPrice.Handle)
}

func configureConsumersMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	groupConsumers := router.Group("/consumer")
	groupConsumers.POST("tn/update/stock", handlers.UpdateStock.Handle)
	groupConsumers.POST("tn/update/price", handlers.UpdatePrice.Handle)
}
