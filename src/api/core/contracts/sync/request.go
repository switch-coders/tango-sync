package sync

import "time"

type Request struct {
	LastUpdate   *time.Time  `form:"last_update" binding:"omitempty"`
}
