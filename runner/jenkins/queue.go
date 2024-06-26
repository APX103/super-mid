package jenkins

import (
	"apx103.com/super-mid/runner/param"
	"github.com/bndr/gojenkins"
)

type JenkinsQueueRunner struct {
	JenkinsClient *gojenkins.Jenkins
}

func NewJenkinsQueueRunner() *JenkinsQueueRunner {
	return &JenkinsQueueRunner{}
}

func (jb *JenkinsQueueRunner) GetCmdPath() string {
	return "jenkins.node"
}

func (jb *JenkinsQueueRunner) GetRunner() func(param.RunnerParam) {
	return func(param param.RunnerParam) {

	}
}
