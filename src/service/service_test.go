package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bancodobrasil/stop-analyzing-api/service/config"
	"github.com/spf13/viper"
	"gotest.tools/assert"
)

func TestUpAndRunning(t *testing.T) {
	builder := new(config.ServiceBuilder).Init(viper.GetViper())
	router := new(Server).InitFromServiceBuilder(builder)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to Stop Analyzing API", w.Body.String())
}
