package jenkins

import "apx103.com/super-mid/command/cmd"

type JenkinsStatusRunner struct {
	JenkinsClient *JenkinsClient
}

func NewJenkinsStatusRunner(jc *JenkinsClient) *JenkinsStatusRunner {
	return &JenkinsStatusRunner{
		JenkinsClient: jc,
	}
}

func (jb *JenkinsStatusRunner) GetCmdPath() string {
	return "jenkins.status"
}

func (jb *JenkinsStatusRunner) Run(*cmd.Task) {
}

func (jb *JenkinsStatusRunner) Finish() {}
