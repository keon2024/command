package factory

import (
	"bufio"
	"command/commands"
	"command/define"
	"fmt"
	"os"
)

const (
	Parse    = "parse"
	Distinct = "distinct"
	Group    = "group"
)

var factory = map[string]define.Command{
	Parse:    &commands.Parse{},
	Distinct: &commands.Distinct{},
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
	flag, result := command.Exec()
	fmt.Println("exec status:", flag)
	// 如果需要打印结果到文件
	argLen := len(args)
	if len(args) < 2 {
		return
	}
	switch args[argLen-2] {
	case "=":
		printToFile(result, args[argLen-1])
	}
	return
}

// CommandDesc 命令介绍
func CommandDesc() {
	for name, command := range factory {
		fmt.Println("command:", name)
		fmt.Println("------------------")
		command.Desc()
	}
	fmt.Println("---------补充---------")
	fmt.Println("命令结尾追加 = filePath 表示将结果打印到文件")
}

// printToFile 打印到文件
func printToFile(result []string, filePath string) {
	if filePath == "" {
		fmt.Println("写出到文件失败: filePath is empty")
		return
	}
	// 打开文件，使用 os.O_CREATE|os.O_WRONLY 表示如果文件不存在就创建，且以写模式打开
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	// 创建一个带缓冲的写入器
	writer := bufio.NewWriter(file)

	// 按行写入文件
	for _, line := range result {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("写入文件失败: %v\n", err)
			return
		}
	}

	// 确保将缓冲区内容刷新到文件
	err = writer.Flush()
	if err != nil {
		fmt.Printf("刷新缓冲区失败: %v\n", err)
		return
	}

	fmt.Println("文件写入成功")
}
