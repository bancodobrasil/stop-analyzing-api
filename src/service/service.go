package service

import (
	"net/http"

	"github.com/bancodobrasil/stop-analyzing-api/db"
	"github.com/bancodobrasil/stop-analyzing-api/service/config"
	v1 "github.com/bancodobrasil/stop-analyzing-api/service/v1"
	ginprom "github.com/banzaicloud/go-gin-prometheus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Server .
type Server struct {
	*config.ServiceBuilder
	app         *gin.Engine
	databaseCli db.DatabasePrisma
	routesV1    v1.Controller
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

	//Configure Database
	s.databaseCli, err = db.Connect()
	if err != nil {
		panic(0)
	}

	return s
}

//RoutesV1 .
func (s *Server) RoutesV1() {
	s.routesV1 = v1.InitRoutesV1(s.databaseCli)

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

	defer s.databaseCli.Disconnect()
	s.app.Run("0.0.0.0:" + s.ServiceBuilder.Port)
}

func index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to Stop Analyzing API")
}
