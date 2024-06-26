package jenkins

import (
	"apx103.com/super-mid/command/cmd"
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

func (jb *JenkinsBuildRunner) GetRunner() func(*cmd.RunnerParamMap) {
	return func(paramMap *cmd.RunnerParamMap) {

	}
}
