package jenkins

import (
	"apx103.com/super-mid/runner/param"
	"github.com/bndr/gojenkins"
)

type JenkinsBuildRunner struct {
	JenkinsClient *gojenkins.Jenkins
}

func NewJenkinsBuildRunner() *JenkinsBuildRunner {
	return &JenkinsBuildRunner{}
}

func (jb *JenkinsBuildRunner) GetCmdPath() string {
	return "jenkins.build"
}

func (jb *JenkinsBuildRunner) GetRunner() func(param.RunnerParam) {
	return func(param param.RunnerParam) {

	}
}
