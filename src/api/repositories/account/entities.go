package account

import (
	"github.com/switch-coders/tango-sync/src/api/core/contracts/integration"
	"strconv"
)

type account struct {
	ID       int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	TangoKey string `gorm:"type:varchar(100);null"`
	Name     string `gorm:"type:varchar(40);null"`
	Email    string `gorm:"type:varchar(100); null"`
	TnKey    string `gorm:"type:varchar(100); null"`
	TnClient int64  `gorm:"type:varchar(100); null"`
}

func newAccountEntity(r integration.Request) *account {
	tnUserID, _ := strconv.Atoi(r.TnUserID)
	return &account{
		Name:     r.Name,
		Email:    r.Email,
		TangoKey: r.TangoKey,
		TnClient: int64(tnUserID),
		TnKey:    r.TnAccessToken,
	}
}
