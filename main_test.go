// server_test.go
package main

import (
	"encoding/json"
	"ip2country/setup"
	"ip2country/store"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	config, err := setup.LoadConfig()
	assert.NoError(t, err)

	gin.SetMode(gin.TestMode)

	db, err := setup.GetDB(config)
	assert.NoError(t, err)

	router := setup.SetupRouter(db, config)
	assert.NotNil(t, router)

	ts := httptest.NewServer(router)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/v1/find-country?ip=8.8.8.8")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var location store.Location
	err = json.NewDecoder(resp.Body).Decode(&location)
	assert.NoError(t, err)
	assert.Equal(t, "United States", location.Country)
	assert.Equal(t, "Mountain View", location.City)
}
