package app

import (
	"github.com/gin-gonic/gin"
	"github.com/switch-coders/tango-sync/src/api/infrastructure/dependencies"
)

func ConfigureMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	configureApiMappings(router, handlers)
	configureJobsMappings(router, handlers)
	configureConsumersMappings(router, handlers)
}

func configureApiMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	group := router.Group("/")
	group.GET("ping", handlers.Get.Handle)
}

func configureJobsMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	groupJobs := router.Group("/jobs")
	groupJobs.POST("tn/sync/stock", handlers.Sync.Handle)
}

func configureConsumersMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	groupConsumers := router.Group("/consumer")
	groupConsumers.POST("tn/sync/stock", handlers.SyncByProduct.Handle)
}
