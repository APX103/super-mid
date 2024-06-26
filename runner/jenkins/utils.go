package jenkins

import (
	"context"

	"github.com/bndr/gojenkins"
)

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
