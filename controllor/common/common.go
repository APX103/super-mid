package common

import "github.com/gin-gonic/gin"

type Controller interface {
	GetPath() string
	GetMethod() string
	GetController() func(c *gin.Context)
}
