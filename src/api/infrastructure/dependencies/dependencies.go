package dependencies

import (
	"github.com/switch-coders/tango-sync/src/api/core/usecases/update_price"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/update_stock"
	"github.com/switch-coders/tango-sync/src/api/repositories/audit"
	"os"

	"github.com/switch-coders/tango-sync/src/api/config/database"
	"github.com/switch-coders/tango-sync/src/api/config/rabbitmq"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/get"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/sync"
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
	Get         entrypoints.Handler
	SyncStock   entrypoints.Handler
	SyncPrice   entrypoints.Handler
	UpdatePrice entrypoints.Handler
	UpdateStock entrypoints.Handler
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

	auditProvider := &audit.Repository{
		DBClient: db,
	}

	// UseCases.
	getUseCase := &get.Implementation{}

	syncUseCase := &sync.Implementation{
		TangoProvider:        tangoProvider,
		ProductProvider:      productProvider,
		NotificationProvider: notificationProvider,
	}

	updateStockUseCase := &update_stock.Implementation{
		TNProvider:      tnProvider,
		ProductProvider: productProvider,
		AuditProvider:   auditProvider,
	}

	updatePriceUseCase := &update_price.Implementation{
		TNProvider:      tnProvider,
		ProductProvider: productProvider,
		AuditProvider:   auditProvider,
	}

	// Handlers.
	handlers := HandlerContainer{}

	handlers.Get = &api.Get{
		GetUseCase: getUseCase,
	}

	handlers.SyncStock = &jobs.SyncStock{
		SyncUseCase: syncUseCase,
	}

	handlers.SyncPrice = &jobs.SyncPrice{
		SyncUseCase: syncUseCase,
	}

	handlers.UpdateStock = &consumer.UpdateStock{
		UpdateStockUseCase: updateStockUseCase,
	}

	handlers.UpdatePrice = &consumer.UpdatePrice{
		UpdatePriceUseCase: updatePriceUseCase,
	}

	return &handlers
}
