package database

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/switch-coders/tango-sync/src/api/config/database/models"
)

func Connect() (Client *gorm.DB, err error) {
	connectionData := GetConnectionDataBase()

	Client, err = gorm.Open("postgres", GetConnectionString(connectionData))

	if err != nil {
		panic(err)
	}

	Client.DB().SetConnMaxLifetime(time.Second * 10)
	Client.DB().SetMaxIdleConns(50)
	Client.DB().SetMaxOpenConns(50)
	Client.SingularTable(true)

	Client.AutoMigrate(&models.Product{}, &models.Audit{}, &models.Account{})

	return
}
