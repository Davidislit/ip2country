package setup

import (
	"errors"
	"ip2country/store"
)

func GetDB(config *Config) (store.DB, error) {
	switch config.DBType {
	case "csv":
		return store.InitCSVDatabase(config.DBPath)
	default:
		return nil, errors.New("not supported DBType")
	}
}
