package commands

import (
	"bufio"
	"fmt"
	"github.com/tidwall/gjson"
	"os"
	"strings"
)

// 支持解析的类型
const (
	json = "json"
)

type Parse struct {
	File string
	Ty   string
	Name string
}

func (p *Parse) Init(args []string) {
	p.File = args[1]
	p.Ty = args[2]
	p.Name = args[3]
	return
}

func (p *Parse) Check(args []string) bool {
	// 参数个数检查
	if len(args) < 4 {
		fmt.Println("params less than 4")
		return false
	}
	// type 类型检查
	switch args[2] {
	case json:
	default:
		fmt.Println("type not support")
		return false
	}
	return true
}

func (p *Parse) Exec() bool {
	fmt.Println("命令执行中", p.File, p.Ty, p.Name)
	switch p.Ty {
	case json:
		jsonParse(p.File, p.Name)
	default:
		fmt.Println("type not support")
		return false
	}
	return true
}

func (p *Parse) Desc() {
	var desc = `parse <file> <type> <name>
	file: 文件路径
	type: 目前只支持json
	name: 字段，支持通过 a.b.c解析多层字段  如果b是一个json字符串，两阶段提取 a.b,c；如果数组则a.0.b.c(如果字段是0则加引号区分)`
	fmt.Println(desc)
}

// json类型文件解析提取字段
func jsonParse(filePath string, name string) {
	var cnt int
	defer func() {
		fmt.Println("总共提取", cnt, "条数据")
	}()
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// 创建一个新的 Reader 对象
	reader := bufio.NewReader(file)

	for {
		// 读取一行
		line, err := reader.ReadString('\n')
		if err != nil {
			// 如果到达文件末尾或遇到错误，则退出循环
			if err.Error() == "EOF" {
				break
			} else {
				fmt.Fprintf(os.Stderr, "读取文件时出错: %v\n", err)
				os.Exit(1)
			}
		}
		names := strings.Split(name, ",")
		var value = line
		for _, n := range names {
			value = gjson.Get(value, n).String()
		}
		if len(value) > 0 {
			cnt++
		}
		fmt.Println(value)

	}
}
