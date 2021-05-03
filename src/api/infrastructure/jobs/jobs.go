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
	tangoSyncProvider := &tango_sync.Repository{
		BaseURL: baseURL,
	}

	s := gocron.NewScheduler(time.UTC)

	job, err := s.Every(2).Hours().Do(tangoSyncProvider.ExecuteSync)
	if err != nil {
		fmt.Println(err)
	}

	job.SingletonMode()
	//s.StartAsync()
}
