package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"apx103.com/super-mid/utils/mongoc"
	"github.com/mattn/go-shellwords"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdParser struct {
	ParsedCmd *TaskCmd
	mongoC    *mongoc.MongoClientImpl
	runnerMap *TaskRunnerMap
	Runner    TaskRunner
}

func NewCmdParser(mongoc *mongoc.MongoClientImpl, rm *TaskRunnerMap) *CmdParser {
	logrus.Debug(" [Fx] CmdParser Init ")
	return &CmdParser{
		mongoC:    mongoc,
		runnerMap: rm,
		Runner:    nil,
		ParsedCmd: nil,
	}
}

func (cp *CmdParser) ParseCommand(cmdString string) string {
	args, err := shellwords.Parse(cmdString)
	if err != nil {
		// TODO 使用MSG模块发送，命令解析错误指令
		fmt.Println(err)
		return ""
	}
	_cmd := cp.CreateParser()
	_cmd.SetArgs(args)
	help_info := CaptureStdout(_cmd.Execute)
	if help_info != "" {
		// TODO 使用MSG模块发送，指令标准输出
		fmt.Println(help_info)
		return help_info
	}
	return ""
}

// Create Parser for each command comes.
func (cp *CmdParser) CreateParser() *cobra.Command {
	temp, ok := cp.mongoC.Find("cmd_template", "", "")

	if !ok {
		fmt.Println("zao!")
		panic("oh ho.")
	}

	var cmdList []*CobraCMD
	for _, t := range *temp {
		group := &CobraCMD{}
		err := json.Unmarshal(t, group)
		if err != nil {
			fmt.Println("unmarshal history error")
			panic("oh ho.")
		}
		cmdList = append(cmdList, group)
	}

	_cmd := &cobra.Command{
		Use:   "@Mr.meeseeks",
		Short: "Feishu Agent Build By QA Team.",
		Long:  `Feishu Agent Build By QA Team. Any question please try '--help' or ask @李佳伦 for help.`,
	}
	ParsedCmd := &TaskCmd{
		CmdPath: "",
		Enable:  true,
		SubCmd:  make(map[string]*TaskCmd),
		Params:  make(map[string]*ParamStruct),
	}
	for _, cmd := range cmdList {
		cp.NewCommand(_cmd, cmd, ParsedCmd)
	}
	cp.ParsedCmd = ParsedCmd
	return _cmd
}

// New command by info come from mongodb `Recursively`.
func (cp *CmdParser) NewCommand(rootCmd *cobra.Command, item *CobraCMD, taskCmd *TaskCmd) {
	_cmdPath := item.CMD
	if taskCmd.CmdPath != "" {
		_cmdPath = taskCmd.CmdPath + "." + _cmdPath
	}
	_taskCmd := &TaskCmd{
		Enable:  false,
		CmdPath: _cmdPath,
		SubCmd:  make(map[string]*TaskCmd),
		Params:  make(map[string]*ParamStruct),
	}
	taskCmd.SubCmd[item.CMD] = _taskCmd

	_cmd := &cobra.Command{
		Use:   item.CMD,
		Short: item.Short,
		Long:  item.Long,
		Run: func(cmd *cobra.Command, args []string) {
			_taskCmd.Enable = true
			// TODO Inject task runner here
			cp.Runner = cp.runnerMap.runners[_cmdPath]
			logrus.Debug("============+ " + _cmdPath + " +============")
		},
	}

	// parse parameters
	if len(item.Params) != 0 {
		for _, param := range item.Params {
			GetCobraType(_cmd, _taskCmd, param)
			if param.Required {
				_cmd.MarkFlagRequired(param.Key)
			}
		}
	}

	// parse sub-command
	if len(item.SubCMD) != 0 {
		for _, subCMD := range item.SubCMD {
			cp.NewCommand(_cmd, &subCMD, _taskCmd)
		}
	}

	rootCmd.AddCommand(_cmd)
}

// Get cobra param type. Add cobra flag type hereif nesessary
func GetCobraType(cmd *cobra.Command, taskCmd *TaskCmd, param CobraParam) {
	switch param.Type {
	case "String":
		taskCmd.Params[param.Key] = &ParamStruct{
			Type:  ParamEnum(param.Type),
			Value: cmd.PersistentFlags().String(param.Key, param.Default, param.Help),
		}
	case "StringToString":
		taskCmd.Params[param.Key] = &ParamStruct{
			Type:  ParamEnum(param.Type),
			Value: cmd.PersistentFlags().StringToString(param.Key, nil, param.Help),
		}
	default:
		fmt.Println("Type not support yet")
	}
}

// Capture cobra.Command.Execute() to get help string
func CaptureStdout(f func() error) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
