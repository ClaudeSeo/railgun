package src

import (
	"github.com/claudeseo/railgun/src/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	pingController := new(controller.PingController)

	router.GET("/ping", pingController.Ping)

	return router
}
