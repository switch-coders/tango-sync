package dependencies

import (
	"os"

	"github.com/switch-coders/tango-sync/src/api/config/database"
	"github.com/switch-coders/tango-sync/src/api/config/rabbitmq"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/get"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/sync"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/sync_by_product"
	"github.com/switch-coders/tango-sync/src/api/entrypoints"
	"github.com/switch-coders/tango-sync/src/api/entrypoints/handlers/api"
	"github.com/switch-coders/tango-sync/src/api/entrypoints/handlers/consumer"
	"github.com/switch-coders/tango-sync/src/api/entrypoints/handlers/jobs"
	"github.com/switch-coders/tango-sync/src/api/repositories/notification"
	"github.com/switch-coders/tango-sync/src/api/repositories/product"
	"github.com/switch-coders/tango-sync/src/api/repositories/tango"
	"github.com/switch-coders/tango-sync/src/api/repositories/tienda_nube"
)

type HandlerContainer struct {
	Get           entrypoints.Handler
	Sync          entrypoints.Handler
	SyncByProduct entrypoints.Handler
}

func Start() *HandlerContainer {
	// Database connection.
	db, err := database.Connect()
	if err != nil {
		panic(errors.ErrorDataBaseConnection)
	}

	ch, err := rabbitmq.Connect()
	if err != nil {
		panic(errors.ErrorConnectingAMQP)
	}

	tangoBaseURL := os.Getenv("TANGO_BASE_URL")
	tangoAccessToken := os.Getenv("TANGO_ACCESS_TOKEN")

	tnBaseURL := os.Getenv("TN_BASE_URL")
	tnAuthentication := os.Getenv("TN_AUTHENTICATION")
	tnUseAgent := os.Getenv("TN_USER_AGENT")
	tnNumber := os.Getenv("TN_NUMBER")

	// Repositories.
	notificationProvider := &notification.Repository{
		Channel: ch,
	}

	tangoProvider := &tango.Repository{
		TangoBaseURL:     tangoBaseURL,
		TangoAccessToken: tangoAccessToken,
	}

	tnProvider := &tienda_nube.Repository{
		TNBaseURL:        tnBaseURL,
		TNAuthentication: tnAuthentication,
		TNUserAgent:      tnUseAgent,
		TNNumber:         tnNumber,
	}

	productProvider := &product.Repository{
		DBClient: db,
	}

	// UseCases.
	getUseCase := &get.Implementation{}

	syncUseCase := &sync.Implementation{
		TangoProvider:        tangoProvider,
		ProductProvider:      productProvider,
		NotificationProvider: notificationProvider,
	}

	syncByProduct := &sync_by_product.Implementation{
		TNProvider:      tnProvider,
		ProductProvider: productProvider,
	}

	// Handlers.
	handlers := HandlerContainer{}

	handlers.Get = &api.Get{
		GetUseCase: getUseCase,
	}

	handlers.Sync = &jobs.Sync{
		SyncUseCase: syncUseCase,
	}

	handlers.SyncByProduct = &consumer.SyncByProduct{
		SyncByProductUseCase: syncByProduct,
	}

	return &handlers
}
