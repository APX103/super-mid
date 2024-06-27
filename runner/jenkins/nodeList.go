package jenkins

import "apx103.com/super-mid/command/cmd"

type JenkinsNodeListRunner struct {
	JenkinsClient *JenkinsClient
}

func NewJenkinsNodeListRunner(jc *JenkinsClient) *JenkinsNodeListRunner {
	return &JenkinsNodeListRunner{
		JenkinsClient: jc,
	}
}

func (jb *JenkinsNodeListRunner) GetCmdPath() string {
	return "jenkins.node.list"
}

func (jb *JenkinsNodeListRunner) Run(*cmd.Task) {
}

func (jb *JenkinsNodeListRunner) Finish() {}
