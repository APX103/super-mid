package jenkins

import (
	"apx103.com/super-mid/command/cmd"
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

func (jb *JenkinsQueueRunner) GetRunner() func(*cmd.RunnerParamMap) {
	return func(paramMap *cmd.RunnerParamMap) {

	}
}
