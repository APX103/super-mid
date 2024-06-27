package jenkins

import (
	"apx103.com/super-mid/command/cmd"
	"github.com/sirupsen/logrus"
)

type JenkinsQueueRunner struct {
	JenkinsClient *JenkinsClient
}

func NewJenkinsQueueRunner(jc *JenkinsClient) *JenkinsQueueRunner {
	return &JenkinsQueueRunner{
		JenkinsClient: jc,
	}
}

func (jb *JenkinsQueueRunner) GetCmdPath() string {
	return "jenkins.node"
}

func (jb *JenkinsQueueRunner) Run(*cmd.Task) {
	logrus.Debug("Run command jenkins queue")
}

func (jb *JenkinsQueueRunner) Finish() {}
