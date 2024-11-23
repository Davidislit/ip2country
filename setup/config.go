package setup

import (
	"os"
	"strconv"
)

type Config struct {
	Port      string
	RateLimit int
	DBType    string
	DBPath    string
}

func LoadConfig() (*Config, error) {
	rateLimitStr := getEnvFromOS("RATE_LIMIT", "5")
	rateLimit, err := strconv.Atoi(rateLimitStr)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:      getEnvFromOS("PORT", "8080"),
		RateLimit: rateLimit,
		DBType:    getEnvFromOS("DB_TYPE", "csv"),
		DBPath:    getEnvFromOS("DB_PATH", "./data/db1.csv"),
	}, nil
}

func getEnvFromOS(key, fallback string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		return fallback
	}

	return value
}
