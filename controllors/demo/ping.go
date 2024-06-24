package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PingController struct {
	Path       string
	Method     string
	Controller func(c *gin.Context)
}

func NewPingController() *PingController {
	return &PingController{}
}

func (c *PingController) GetPath() string {
	return "/ping"
}

func (c *PingController) GetMethod() string {
	return "GET"
}

func (c *PingController) GetController() func(c *gin.Context) {
	return func(c *gin.Context) {
		logrus.Info("Heartbeat. Health Check")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
