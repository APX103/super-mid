package jenkins

import (
	"context"

	"apx103.com/super-mid/utils/config"
	"github.com/bndr/gojenkins"
	"github.com/sirupsen/logrus"
)

type JenkinsClient struct {
	ctx context.Context
	C   *gojenkins.Jenkins
}

func NewJenkinsClient(conf *config.BaseConfig) *JenkinsClient {
	logrus.Debug(" [Fx] JenkinsClient Init ")
	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, conf.JenkinsConfig.Url, conf.JenkinsConfig.Username, conf.JenkinsConfig.Password)
	return &JenkinsClient{
		ctx: ctx,
		C:   jenkins,
	}
}

func (jc *JenkinsClient) Init() {
	jc.C.Init(jc.ctx)
}

func GetAllJobs(ctx context.Context, jenkins *gojenkins.Jenkins, jobs []*gojenkins.Job, parents string) []*JobStruct {
	jobNames := []*JobStruct{}
	for _, job := range jobs {
		_parents := parents
		if job.Raw.Class == "com.cloudbees.hudson.plugins.folder.Folder" {
			if _parents == "" {
				_parents = job.Raw.Name
			} else {
				_parents += "/" + job.Raw.Name
			}
			subJobs, _ := job.GetInnerJobs(ctx)
			subJobNames := GetAllJobs(ctx, jenkins, subJobs, _parents)
			jobNames = append(jobNames, subJobNames...)
		} else {
			jobNames = append(jobNames, &JobStruct{
				Job:     job.GetName(),
				Parents: _parents,
			})
		}
	}
	return jobNames
}
