package cmd

import (
	"github.com/sirupsen/logrus"
)

type TaskRunnerMap struct {
	runners map[string]TaskRunner
}

func NewCmdRunnerMap(runners []TaskRunner) *TaskRunnerMap {
	logrus.Debug(" [Fx] TaskRunnerMap Init ")
	_map := &TaskRunnerMap{
		runners: make(map[string]TaskRunner),
	}

	for _, runner := range runners {
		_map.runners[runner.GetCmdPath()] = runner
	}
	return _map
}
