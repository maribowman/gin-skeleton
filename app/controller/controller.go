package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maribowman/gin-skeleton/app/config"
	"github.com/maribowman/gin-skeleton/app/controller/middleware"
	"github.com/maribowman/gin-skeleton/app/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Controller struct {
	router            *gin.Engine
	service           model.Service
	prometheusHandler http.Handler
}

type ControllerWiring struct {
	Router            *gin.Engine
	Service           model.Service
	PrometheusHandler http.Handler
}

func NewController(wiring *ControllerWiring) {
	controller := &Controller{
		router:            wiring.Router,
		service:           wiring.Service,
		prometheusHandler: wiring.PrometheusHandler,
	}
	controller.router.Use(gin.Logger(), gin.Recovery())

	log.Info().Bool("hmacEnabled", config.Config.Authentication.HmacEnabled).Msg("TEST")

	protected := controller.router.Group("/protected").
		Use(middleware.HmacMiddleware(config.Config.Authentication.HmacEnabled, config.Config.Authentication.AllowedUsers))

	protected.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "(┛❍ᴥ❍)┛"})
		return
	})

	controller.router.GET("/metrics", func(c *gin.Context) {
		controller.prometheusHandler.ServeHTTP(c.Writer, c.Request)
	})
}
