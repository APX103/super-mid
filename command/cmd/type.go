package cmd

// Define cobra parser struct
// Cobra parameters struct(json)
type CobraParam struct {
	Key      string `json:"key" bson:"key"`
	Type     string `json:"type" bson:"type"`
	Default  string `json:"default,omitempty" bson:"default,omitempty"`
	Required bool   `json:"required" bson:"required"`
	Help     string `json:"help,omitempty" bson:"help,omitempty"`
}

// Cobra command struct(json)
type CobraCMD struct {
	CMD    string       `json:"cmd" bson:"cmd"`
	SubCMD []CobraCMD   `json:"sub-cmd,omitempty" bson:"cmd,omitempty"`
	Params []CobraParam `json:"params,omitempty" bson:"params,omitempty"`
	Short  string       `json:"short,omitempty" bson:"short,omitempty"`
	Long   string       `json:"long,omitempty" bson:"long,omitempty"`
}

// Define type enumerate for task param
type ParamEnum string

// only support 2 types of cobra flag type yet(S and S2S).
const String ParamEnum = "String"
const Int ParamEnum = "Int"
const Bool ParamEnum = "Bool"
const StringToString ParamEnum = "StringToString"
const StringToInt ParamEnum = "StringToInt"
const StringToBool ParamEnum = "StringToBool"

// Define param struct(multi type input param)
type ParamStruct struct {
	Type  ParamEnum
	Value any
}

// Task cmd struct for transfer cmd and param to TaskInfo
type TaskCmd struct {
	CmdPath string
	SubCmd  map[string]*TaskCmd
	Enable  bool
	Params  map[string]*ParamStruct
}

// Task runner interface define.
type RunnerParam interface {
	GetCmdPath() string
	GetParamStruct()
}

type RunnerParamMap struct {
	Map map[string]RunnerParam
}

type TaskRunner interface {
	GetCmdPath() string
	GetRunner() func(paramMap *RunnerParamMap)
}

func NewRunnerParamMap(runnerParams []RunnerParam) *RunnerParamMap {
	_map := &RunnerParamMap{
		Map: make(map[string]RunnerParam),
	}
	for _, rp := range runnerParams {
		_map.Map[rp.GetCmdPath()] = rp
	}
	return _map
}
