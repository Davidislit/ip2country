package api

import (
	"net"
	"net/http"

	"ip2country/store"

	"github.com/gin-gonic/gin"
)

type FindCountryService struct {
	db store.DB
}

func RegisterFindCountryRoute(router *gin.RouterGroup, db store.DB) {

	findCountryService := FindCountryService{db: db}

	router.GET("find-country", findCountryService.findCountry)

}

func (service *FindCountryService) findCountry(c *gin.Context) {
	ip := c.Query("ip")

	if ip == "" {
		c.JSON(http.StatusBadRequest, ResponseError{Error: "ip query parameter is required"})
		return
	}

	if net.ParseIP(ip) == nil {
		c.JSON(http.StatusBadRequest, ResponseError{Error: "invalid ip address"})
		return
	}

	location, err := service.db.Find(ip)
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseError{Error: "Not found ip"})
		return
	}

	c.JSON(http.StatusOK, location)
}
