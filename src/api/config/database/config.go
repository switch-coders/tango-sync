package database

import (
	"fmt"
	"os"
	"strings"
)

type ConnectionData struct {
	Host     string
	Schema   string
	Username string
	Password string
	SSL      string
}

func GetConnectionDataBase() *ConnectionData {
	scope := os.Getenv("SCOPE")
	connectionData := ConnectionData{}
	if strings.HasSuffix(scope, "master") {
		return connectionData.setupMasterConnectionData()
	}

	if strings.HasSuffix(scope, "beta") {
		return connectionData.setupMasterConnectionData()
	}

	connectionData.Host = "localhost"
	connectionData.Schema = "tango-sync"
	connectionData.Username = "nmaltagliatt"
	connectionData.Password = "Menani437."
	connectionData.SSL = "disable"
	return &connectionData
}

func (cd *ConnectionData) setupMasterConnectionData() *ConnectionData {
	cd.Host = "ec2-54-224-120-186.compute-1.amazonaws.com"
	cd.Password = "085e0729864a3343abf08af4f27e67c93d1a73ab724053b35914f087d5c207ca"
	cd.Username = "qsdidhxtocenoh"
	cd.Schema = "dcqrc5tjn0av78"
	cd.SSL = "require"
	return cd
}

func (cd *ConnectionData) setupBetaConnectionData() *ConnectionData {
	cd.Host = "ec2-3-91-127-228.compute-1.amazonaws.com"
	cd.Password = "ecf48c2caafeacbbe259f68d49bbebf6e0de4df3d10eeb95a6f52529935d6f27"
	cd.Username = "zkgvkqrdpvtrix"
	cd.Schema = "df5ica15dgfqbo"
	cd.SSL = "require"
	return cd
}

func GetConnectionString(cd *ConnectionData) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s", cd.Host, cd.Username, cd.Password, cd.Schema, cd.SSL)
}
