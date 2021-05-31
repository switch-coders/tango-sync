package account

import "github.com/switch-coders/tango-sync/src/api/core/contracts/integration"

type account struct {
	TangoKey string `gorm:"PRIMARY_KEY;type:varchar(100);not null"`
	Name     string `gorm:"type:varchar(40);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	JobSync  bool
	JobPrice bool
}

func newAccountEntity(r integration.Request) *account {
	return &account{
		Name:     r.Name,
		Email:    r.Email,
		TangoKey: r.TangoKey,
		JobSync:  true,
		JobPrice: true,
	}
}
