package routes

import (
	"net/http"

	"github.com/claudeseo/railgun/controllers"
	"github.com/gin-gonic/gin"
)

func globalRecovery(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "서버에 문제가 발생했습니다.",
				"status":  http.StatusInternalServerError,
			})
		}
	}(c)
	c.Next()
}

func notFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "잘못된 API 경로입니다.",
		"status":  http.StatusNotFound,
	})
}

func Init() *gin.Engine {
	pingController := new(controllers.PingController)
	placeController := new(controllers.PlaceController)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(globalRecovery)
	router.NoRoute(notFoundHandler)
	router.GET("/ping", pingController.Ping)
	v1 := router.Group("/v1")
	{
		v1Place := v1.Group("/places")
		{
			v1Place.GET("", placeController.GetPlace)
		}
	}

	return router
}
