package service

import (
	"net/http"

	"github.com/bancodobrasil/stop-analyzing-api/service/config"
	ginprom "github.com/banzaicloud/go-gin-prometheus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Server .
type Server struct {
	*config.ServiceBuilder
	app *gin.Engine
}

//InitFromServiceBuilder builds a Server instance
func (s *Server) InitFromServiceBuilder(serviceBuilder *config.ServiceBuilder) *Server {
	s.ServiceBuilder = serviceBuilder
	s.app = gin.New()

	logLevel, err := logrus.ParseLevel(s.ServiceBuilder.LogLevel)
	if err != nil {
		logrus.Errorf("Not able to parse log level string. Setting default level: info.")
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)

	p := ginprom.NewPrometheus("gin", []string{})
	p.Use(s.app, "/metrics")

	//Configure Routes
	s.Routes()

	return s
}

//Routes .
func (s *Server) Routes() {
	s.app.GET("/", index)

	// Versions
	v1Group := s.app.Group("/v1/")
	{
		v1Group.GET("/", index)
		// consumerGroup.GET("/owner/:owner/thing/:thing/node/:node", s.consumer.GetHandler)
		// consumerGroup.POST("/owner/:owner/thing/:thing/node/:node", s.consumer.CreateHandler)
	}
}

//Run starts the http server service
func (s *Server) Run() {
	logrus.Info("Version 0.0.1")
	s.app.Run("0.0.0.0:" + s.ServiceBuilder.Port)
}

func index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to Stop Analyzing API")
}
