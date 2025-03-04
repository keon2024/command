package commands

import (
	"command/utils"
	"fmt"
	"strings"
)

type Distinct struct {
	File string
}

func (d *Distinct) Init(args []string) {
	d.File = args[1]
}

func (d *Distinct) Check(args []string) bool {
	// 参数个数检查
	if len(args) < 2 {
		fmt.Println("params less than 2")
		return false
	}
	return true
}

func (d *Distinct) Exec() (bool, []string) {
	fmt.Println("命令执行中", d.File)
	var (
		result []string
		flag   bool
	)

	flag, result = readFileDistinct(d.File)

	return flag, result
}

func (d *Distinct) Desc() {
	var desc = `distinct <file> 
	file: 文件路径`
	fmt.Println(desc)
}

// readFileDistinct 读取文件并去重
func readFileDistinct(filePath string) (
	bool, []string) {
	var (
		flag   bool
		result []string
		dm     = make(map[string]int8)
	)
	flag = utils.ReadFile(filePath, func(line string) bool {
		value := strings.TrimSpace(line)
		if _, ok := dm[value]; !ok && value != "" {
			dm[value] = 1
			result = append(result, value)
			fmt.Println(value)
			return true
		}
		return false
	})

	return flag, result
}
