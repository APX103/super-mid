package jenkins

import (
	"apx103.com/super-mid/command/cmd"
	"github.com/sirupsen/logrus"
)

type JenkinsStatusRunner struct {
	JenkinsClient *JenkinsClient
}

func NewJenkinsStatusRunner(jc *JenkinsClient) *JenkinsStatusRunner {
	return &JenkinsStatusRunner{
		JenkinsClient: jc,
	}
}

func (jb *JenkinsStatusRunner) GetCmdPath() string {
	return "jenkins.status"
}

func (jb *JenkinsStatusRunner) Run(*cmd.Task) {
	logrus.Debug("Run command jenkins status")
}

func (jb *JenkinsStatusRunner) Finish() {}
