package commands

import (
	"command/utils"
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

// 支持解析的类型
const (
	json = "json"
	form = "form" // 表单类型，日志中展示为 opt_time=123&order_id=456 格式
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
	case form:
	default:
		fmt.Println("type not support")
		return false
	}
	return true
}

func (p *Parse) Exec() (bool, []string) {
	fmt.Println("命令执行中", p.File, p.Ty, p.Name)
	var (
		result []string
		flag   bool
	)
	switch p.Ty {
	case json:
		flag, result = jsonParse(p.File, p.Name)
	case form:
		flag, result = formParse(p.File, p.Name)
	default:
		fmt.Println("type not support")
		flag = false
	}
	return flag, result
}

func (p *Parse) Desc() {
	var desc = `parse <file> <type> <name>
	file: 文件路径
	type: 目前只支持json、form
	name: 字段
		json支持通过 a.b.c解析多层字段  如果b是一个json字符串，两阶段提取 a.b,c；如果数组则a.0.b.c(如果字段是0则加引号区分)
		form直接通过名字提取，如 opt_time=123&order_id=456 格式 name=opt_time`
	fmt.Println(desc)
}

// jsonParse json类型文件解析提取字段
func jsonParse(filePath string, name string) (bool, []string) {
	var (
		flag   bool
		result []string
	)
	flag = utils.ReadFile(filePath, func(line string) bool {
		names := strings.Split(name, ",")
		var value = line
		for _, n := range names {
			value = gjson.Get(value, n).String()
		}
		fmt.Println(value)
		if value != "" {
			result = append(result, value)
			return true
		}
		return false
	})

	return flag, result
}

// formParse form类型字段提取 opt_time=123&order_id=456 格式
func formParse(filePath string, name string) (bool, []string) {
	var (
		flag   bool
		result []string
	)
	flag = utils.ReadFile(filePath, func(line string) bool {
		line = strings.TrimSpace(line)
		vs := strings.Split(line, "&")
		var value string
		for _, v := range vs {
			start := strings.Index(v, name+"=")
			if start != 0 {
				continue
			}
			end := len(name + "=")
			value = v[end:]
			break
		}
		fmt.Println(value)
		if value != "" {
			result = append(result, value)
			return true
		}
		return false
	})

	return flag, result
}
