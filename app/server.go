package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maribowman/gin-skeleton/app/config"
	"github.com/maribowman/gin-skeleton/app/controller"
	"github.com/maribowman/gin-skeleton/app/model"
	"github.com/maribowman/gin-skeleton/app/repository"
	"github.com/maribowman/gin-skeleton/app/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func InitServer(databaseClient model.DatabaseClient) (*http.Server, error) {
	gin.SetMode(config.Config.Server.Mode)
	router := gin.New()
	controller.NewController(&controller.ControllerWiring{
		Router: router,
		Service: service.NewService(service.ServiceWiring{
			DatabaseClient: databaseClient,
			RestClient:     repository.NewDemoRestClient(),
		}),
		PrometheusHandler: promhttp.Handler(),
	})

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Config.Server.Port),
		Handler: router,
	}, nil
}
