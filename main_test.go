// server_test.go
package main

import (
	"encoding/json"
	"ip2country/api"
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

	respIpIsRequired, err := http.Get(ts.URL + "/v1/find-country")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, respIpIsRequired.StatusCode)

	var resposneIpIsRequired api.ResponseError
	err = json.NewDecoder(respIpIsRequired.Body).Decode(&resposneIpIsRequired)
	assert.NoError(t, err)
	assert.Equal(t, "ip query parameter is required", resposneIpIsRequired.Error)

	respInvalidIp, err := http.Get(ts.URL + "/v1/find-country?ip=8.8.8.")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, respInvalidIp.StatusCode)

	var responseInvalidIp api.ResponseError
	err = json.NewDecoder(respInvalidIp.Body).Decode(&responseInvalidIp)
	assert.NoError(t, err)
	assert.Equal(t, "invalid ip address", responseInvalidIp.Error)

	respNotFoundIp, err := http.Get(ts.URL + "/v1/find-country?ip=8.8.8.0")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, respNotFoundIp.StatusCode)

	var responseNotFoundIp api.ResponseError
	err = json.NewDecoder(respNotFoundIp.Body).Decode(&responseNotFoundIp)
	assert.NoError(t, err)
	assert.Equal(t, "Not found ip", responseNotFoundIp.Error)

}
