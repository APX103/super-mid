package jenkins

import (
	"apx103.com/super-mid/runner/param"
	"github.com/bndr/gojenkins"
)

type JenkinsNodeRunner struct {
	JenkinsClient *gojenkins.Jenkins
}

func NewJenkinsNodedRunner() *JenkinsNodeRunner {
	return &JenkinsNodeRunner{}
}

func (jb *JenkinsNodeRunner) GetCmdPath() string {
	return "jenkins.node"
}

func (jb *JenkinsNodeRunner) GetRunner() func(param.RunnerParam) {
	return func(param param.RunnerParam) {

	}
}
