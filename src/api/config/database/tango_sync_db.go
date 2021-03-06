package database

import (
	"github.com/switch-coders/tango-sync/src/api/config/database/models"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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

	Client.AutoMigrate(&models.Product{}, &models.Audit{})

	return
}
