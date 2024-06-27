package jenkins

import (
	"apx103.com/super-mid/command/cmd"
	"github.com/sirupsen/logrus"
)

type JenkinsBuildRunner struct {
	JenkinsClient *JenkinsClient
}

func NewJenkinsBuildRunner(jc *JenkinsClient) *JenkinsBuildRunner {
	return &JenkinsBuildRunner{
		JenkinsClient: jc,
	}
}

func (jb *JenkinsBuildRunner) GetCmdPath() string {
	return "jenkins.build"
}

func (jb *JenkinsBuildRunner) Run(*cmd.Task) {
	logrus.Debug("Run command jenkins build")
}

func (jb *JenkinsBuildRunner) Finish() {}
