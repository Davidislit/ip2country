package setup

import (
	"ip2country/api"
	"ip2country/middleware"
	"ip2country/store"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db store.DB, config *Config) *gin.Engine {
	router := gin.Default()

	api.RegisterHealthRoute(router)

	rateLimiter := middleware.NewRateLimiter(config.RateLimit)

	v1 := router.Group("/v1")

	v1.Use(middleware.RateLimitMiddleware(rateLimiter))

	api.RegisterFindCountryRoute(v1, db)

	return router
}
