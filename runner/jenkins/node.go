package jenkins

import (
	"apx103.com/super-mid/command/cmd"
	"github.com/sirupsen/logrus"
)

type JenkinsNodeRunner struct {
	JenkinsClient *JenkinsClient
}

func NewJenkinsNodeRunner(jc *JenkinsClient) *JenkinsNodeRunner {
	return &JenkinsNodeRunner{
		JenkinsClient: jc,
	}
}

func (jb *JenkinsNodeRunner) GetCmdPath() string {
	return "jenkins.node"
}

func (jb *JenkinsNodeRunner) Run(*cmd.Task) {
	logrus.Debug("Run command jenkins node")
}

func (jb *JenkinsNodeRunner) Finish() {}
