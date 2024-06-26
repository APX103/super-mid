package param

type JenkinsBuildParam struct {
	JobName string
	Params  map[string]string
}

type JenkinsListParam struct {
	Folder string
}

type JenkinsQueueParam struct {
}

type RunnerParam struct {
	Build JenkinsBuildParam
	List  JenkinsListParam
	Queue JenkinsQueueParam
}
