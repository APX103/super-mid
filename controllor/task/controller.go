package task

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TaskController struct {
	Path       string
	Method     string
	Controller func(c *gin.Context)
}

func NewPingController() *TaskController {
	controller := &TaskController{
		Path:   "/task",
		Method: "POST",
	}

	controller.Controller = func(c *gin.Context) {
		logrus.Info("web post req /api/config")
		req := &WebPostTaskRequest{}
		bytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			logrus.WithError(err).Errorf("failed to read WebPostTaskRequest")
			return
		}

		err = json.Unmarshal(bytes, req)
		if err != nil {
			logrus.WithError(err).Errorf("failed to unmarsal WebPostTaskRequest")
			c.JSON(500, gin.H{
				"message": "Wrong ",
			})
			return
		}
	}
	return controller
}

func (tc *TaskController) GetPath() string {
	return tc.Path
}

func (tc *TaskController) GetMethod() string {
	return tc.Method
}

func (tc *TaskController) GetController() func(c *gin.Context) {
	return func(c *gin.Context) {
		logrus.Info("Heartbeat. Health Check")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
