package jenkins

import (
	"apx103.com/super-mid/runner/param"
	"github.com/bndr/gojenkins"
)

type JenkinsListRunner struct {
	JenkinsClient *gojenkins.Jenkins
}

func NewJenkinsListdRunner() *JenkinsListRunner {
	return &JenkinsListRunner{}
}

func (jb *JenkinsListRunner) GetCmdPath() string {
	return "jenkins.list"
}

func (jb *JenkinsListRunner) GetRunner() func(param.RunnerParam) {
	return func(param param.RunnerParam) {

	}
}
