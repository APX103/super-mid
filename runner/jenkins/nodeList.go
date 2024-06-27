package jenkins

import (
	"apx103.com/super-mid/command/cmd"
	"github.com/sirupsen/logrus"
)

type JenkinsNodeListRunner struct {
	JenkinsClient *JenkinsClient
}

func NewJenkinsNodeListRunner(jc *JenkinsClient) *JenkinsNodeListRunner {
	return &JenkinsNodeListRunner{
		JenkinsClient: jc,
	}
}

func (jb *JenkinsNodeListRunner) GetCmdPath() string {
	return "jenkins.node.list"
}

func (jb *JenkinsNodeListRunner) Run(*cmd.Task) {
	logrus.Debug("Run command jenkins node list")
	// 获取job是否存在
	// 获取job param
	// 验证入参
	// Jenkins build job
}

func (jb *JenkinsNodeListRunner) Finish() {}
