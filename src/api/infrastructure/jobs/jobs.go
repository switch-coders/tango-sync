package jobs

import (
	"fmt"
	"os"
	"time"

	"github.com/go-co-op/gocron"

	"github.com/switch-coders/tango-sync/src/api/repositories/tango_sync"
)

func Schedule() {
	baseURL := os.Getenv("BASE_URL")

	// Repositories.
	emcSyncProvider := &tango_sync.Repository{
		BaseURL: baseURL,
	}

	s := gocron.NewScheduler(time.UTC)

	now := time.Now().UTC()
	stockStartDate := now.Add(time.Minute * 30)
	//priceStartDate := now.Add(time.Minute * 90)

	jobSyncStock, err := s.Every(2).Hours().StartAt(stockStartDate).Do(emcSyncProvider.ExecuteStockSync)
	if err != nil {
		fmt.Println(err)
	}

	/*	jobSyncPrice, err := s.Every(2).Hours().StartAt(priceStartDate).Do(emcSyncProvider.ExecutePriceSync)
		if err != nil {
			fmt.Println(err)
		}*/

	jobSyncStock.SingletonMode()
	//jobSyncPrice.SingletonMode()

	s.StartAsync()
}
