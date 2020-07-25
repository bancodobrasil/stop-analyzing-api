package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bancodobrasil/stop-analyzing-api/internal/api/config"
	"github.com/spf13/viper"
	"gotest.tools/assert"
)

func TestUpAndRunning(t *testing.T) {
	builder := new(config.APIBuilder).Init(viper.GetViper())
	router := new(Server).InitFromAPIBuilder(builder)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to Stop Analyzing API", w.Body.String())
}

func TestRegisterRoutesV1(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://user2020:pass2020@localhost:5432/stop-analyzing-api")

	builder := new(config.APIBuilder).Init(viper.GetViper())
	router := new(Server).InitFromAPIBuilder(builder)
	//Register V1 Routes
	router.RoutesV1()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/", nil)
	router.app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Version 1 API", w.Body.String())
}
