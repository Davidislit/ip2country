package setup

import (
	"errors"
	"fmt"
	"ip2country/store"
)

func GetDB(config *Config) (store.DB, error) {
	switch config.DBType {
	case "csv":
		return store.InitCSVDatabase(config.DBPath)
	default:
		return nil, errors.New(fmt.Sprint("not supported DBType %v", config.DBType))
	}
}
