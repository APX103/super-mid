package jenkins

import (
	"apx103.com/super-mid/command/cmd"
	"github.com/sirupsen/logrus"
)

type JenkinsNodeOnlineRunner struct {
	JenkinsClient *JenkinsClient
}

func NewJenkinsNodeOnlineRunner(jc *JenkinsClient) *JenkinsNodeOnlineRunner {
	return &JenkinsNodeOnlineRunner{
		JenkinsClient: jc,
	}
}

func (jb *JenkinsNodeOnlineRunner) GetCmdPath() string {
	return "jenkins.node.online"
}

func (jb *JenkinsNodeOnlineRunner) Run(*cmd.Task) {
	logrus.Debug("Run command jenkins online")
}

func (jb *JenkinsNodeOnlineRunner) Finish() {}
