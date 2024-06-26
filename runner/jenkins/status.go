package jenkins

import (
	"apx103.com/super-mid/runner/param"
	"github.com/bndr/gojenkins"
)

type JenkinsStatusRunner struct {
	JenkinsClient *gojenkins.Jenkins
}

func NewJenkinsStatusRunner() *JenkinsStatusRunner {
	return &JenkinsStatusRunner{}
}

func (jb *JenkinsStatusRunner) GetCmdPath() string {
	return "jenkins.node"
}

func (jb *JenkinsStatusRunner) GetRunner() func(param.RunnerParam) {
	return func(param param.RunnerParam) {

	}
}
