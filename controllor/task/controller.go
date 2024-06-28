package task

import (
	"encoding/json"
	"io"
	"strings"

	"apx103.com/super-mid/command/cmd"
	"apx103.com/super-mid/message/feishu"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type TaskController struct {
	Path          string
	Method        string
	Controller    func(c *gin.Context)
	CommandParser *cmd.CmdParser
}

func NewTaskController(cp *cmd.CmdParser) *TaskController {
	logrus.Debug(" [Fx] TaskController Init ")
	controller := &TaskController{
		Path:          "/task",
		Method:        "POST",
		CommandParser: cp,
	}

	controller.Controller = func(c *gin.Context) {
		logrus.Info("web post req /task")
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
		logrus.Infof("Command line is: %s", req.CommandLine)
		// cmd parse
		cmdStr := controller.CommandParser.ParseCommand(req.CommandLine)
		if strings.Contains(cmdStr, "help") {
			resp := &RunnerWebResponse{
				Code:       "0",
				Message:    "",
				Content:    feishu.BuildSimpleFeishuCardWithText("Help", cmdStr, "green"),
				HasContent: true,
			}
			c.JSON(200, resp)
			return
		} else if cmdStr != "" {
			resp := &RunnerWebResponse{
				Code:       "1",
				Message:    "",
				Content:    feishu.BuildSimpleFeishuCardWithText("Error", cmdStr, "red"),
				HasContent: true,
			}
			c.JSON(200, resp)
			return
		}
		taskInfo := &cmd.Task{
			TaskID:           uuid.NewV1().String(),
			User:             req.UserName,
			UserFeishuID:     req.FeishuSenderID,
			UserFeishuIDType: req.FeishuSenderIDType,
			FeishuChatID:     req.FeishuChatID,
			TaskCmd:          controller.CommandParser.ParsedCmd,
		}

		// runner
		if controller.CommandParser.Runner != nil {
			logrus.Debugf("Run path: %s", controller.CommandParser.Runner.GetCmdPath())
			controller.CommandParser.Runner.Run(taskInfo)
			resp := &RunnerWebResponse{
				Code:       "0",
				Message:    "",
				Content:    feishu.BuildSimpleFeishuCardWithText("Task Run", controller.CommandParser.Runner.GetCmdPath()+" task launched.", "green"),
				HasContent: true,
			}
			c.JSON(200, resp)
			return
		}
		// TODO Push taskInfo to Task Queue with CmdPath
		logrus.Debugf("CMD path: %s. Pushed to TaskQueue", controller.CommandParser.CmdPath)
		resp := &RunnerWebResponse{
			Code:       "0",
			Message:    "",
			Content:    feishu.BuildSimpleFeishuCardWithText("Task Push", controller.CommandParser.Runner.GetCmdPath()+" task pushed.(no runner inside)", "green"),
			HasContent: true,
		}
		c.JSON(200, resp)
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
	return tc.Controller
}
