package api

import (
	v12 "github.com/bancodobrasil/stop-analyzing-api/internal/api/v1"
	"github.com/bancodobrasil/stop-analyzing-api/internal/domain"
	"net/http"

	"github.com/bancodobrasil/stop-analyzing-api/internal/api/config"
	ginprom "github.com/banzaicloud/go-gin-prometheus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*config.APIBuilder
	app      *gin.Engine
	service  *domain.Service
	routesV1 v12.Controller
}

//InitFromAPIBuilder builds a Server instance
func (s *Server) InitFromAPIBuilder(serviceBuilder *config.APIBuilder) *Server {
	s.APIBuilder = serviceBuilder
	s.app = gin.New()

	logLevel, err := logrus.ParseLevel(s.APIBuilder.LogLevel)
	if err != nil {
		logrus.Errorf("Not able to parse log level string. Setting default level: info.")
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)

	p := ginprom.NewPrometheus("gin", []string{})
	p.Use(s.app, "/metrics")

	//Configure Database
	s.service, err = domain.NewService()
	if err != nil {
		panic(0)
	}

	return s
}

//RoutesV1 .
func (s *Server) RoutesV1() {
	s.routesV1 = v12.InitRoutesV1(s.service)

	v1Group := s.app.Group("/v1/")
	{
		v1Group.GET("/", s.routesV1.Index)
		v1Group.GET("/listTags", s.routesV1.ListAllTags)
	}
}

//Routes .
func (s *Server) Routes() {
	s.app.GET("/", index)
}

//Run starts the http server service
func (s *Server) Run() {
	logrus.Info("Version 0.0.1")

	//Configure Routes
	s.Routes()
	s.RoutesV1()

	defer s.service.Disconnect()
	s.app.Run("0.0.0.0:" + s.APIBuilder.Port)
}

func index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to Stop Analyzing API")
}
