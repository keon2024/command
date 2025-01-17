package factory

import (
	"command/commands"
	"command/define"
	"fmt"
)

const (
	Parse = "parse"
)

var factory = map[string]define.Command{
	Parse: &commands.Parse{},
}

// ExecCommand 执行命令
func ExecCommand(name string, args []string) {
	command, ok := factory[name]
	if !ok {
		fmt.Println("command not found")
		return
	}
	if !command.Check(args) {
		fmt.Println("command check failed")
		return
	}
	command.Init(args)
	result := command.Exec()
	fmt.Println("exec result:", result)
	return
}

// CommandDesc 命令介绍
func CommandDesc() {
	for name, command := range factory {
		fmt.Println("command:", name)
		fmt.Println("------------------")
		command.Desc()
	}
}
