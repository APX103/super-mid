package jenkins

import (
	"apx103.com/super-mid/command/cmd"
)

type JenkinsNodeRunner struct {
	JenkinsClient *JenkinsClient
}

func NewJenkinsNodedRunner(jc *JenkinsClient) *JenkinsNodeRunner {
	return &JenkinsNodeRunner{
		JenkinsClient: jc,
	}
}

func (jb *JenkinsNodeRunner) GetCmdPath() string {
	return "jenkins.node"
}

func (jb *JenkinsNodeRunner) GetRunner() func(*cmd.RunnerParamMap) {
	return func(paramMap *cmd.RunnerParamMap) {

	}
}
