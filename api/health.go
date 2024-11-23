package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoute(router *gin.Engine) {
	router.GET("/health", healthCheck)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
