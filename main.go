package main

import (
	"ip2country/setup"
	"log"
)

func bootstrapServer() {
	log.Printf("Starting ip2country server")

	config, err := setup.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading config %v", err)
	}

	db, err := setup.GetDB(config)
	if err != nil {
		log.Fatalf("Error initalizing database: %v", err)
	}

	router := setup.SetupRouter(db, config)

	log.Printf("Starting ip2country server on port: %v", config.Port)

	err = router.Run(":" + config.Port)

	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}

func main() {
	bootstrapServer()
}
