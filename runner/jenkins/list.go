package jenkins

import (
	"apx103.com/super-mid/command/cmd"
)

type JenkinsListRunner struct {
	JenkinsClient *JenkinsClient
}

func NewJenkinsListdRunner(jc *JenkinsClient) *JenkinsListRunner {
	return &JenkinsListRunner{
		JenkinsClient: jc,
	}
}

func (jb *JenkinsListRunner) GetCmdPath() string {
	return "jenkins.list"
}

func (jb *JenkinsListRunner) GetRunner() func(*cmd.RunnerParamMap) {
	return func(paramMap *cmd.RunnerParamMap) {

	}
}
