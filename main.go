package main

import (
	"command/factory"
	"flag"
	"fmt"
)

/*
命令行定义
第一个字段表示命令
然后根据不同命令后面字段有不同定义和实现
*/
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return
	}
	if args[0] == "help" {
		fmt.Println("Usage: main <command>")
		factory.CommandDesc()
		return
	}
	// 执行命令
	factory.ExecCommand(args[0], args)

}
