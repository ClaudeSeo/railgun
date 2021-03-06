package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (ctrl PingController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PONG",
	})
}
