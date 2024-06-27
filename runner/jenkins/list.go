package jenkins

import "apx103.com/super-mid/command/cmd"

type JenkinsListRunner struct {
	JenkinsClient *JenkinsClient
}

func NewJenkinsListRunner(jc *JenkinsClient) *JenkinsListRunner {
	return &JenkinsListRunner{
		JenkinsClient: jc,
	}
}

func (jb *JenkinsListRunner) GetCmdPath() string {
	return "jenkins.list"
}

func (jb *JenkinsListRunner) Run(*cmd.Task) {
}

func (jb *JenkinsListRunner) Finish() {}
