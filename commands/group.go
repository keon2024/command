package commands

import (
	"command/utils"
	"fmt"
	"strings"
)

type Group struct {
	File string
}

func (g *Group) Init(args []string) {
	g.File = args[1]
}

func (g *Group) Check(args []string) bool {
	// 参数个数检查
	if len(args) < 2 {
		fmt.Println("params less than 2")
		return false
	}
	return true
}

func (g *Group) Exec() (bool, []string) {
	fmt.Println("命令执行中", g.File)
	var (
		result []string
		flag   bool
	)

	flag, result = readFileGroup(g.File)

	return flag, result
}

func (g *Group) Desc() {
	var desc = `group <file> 
	file: 文件路径`
	fmt.Println(desc)
}

// readFileGroup 读取文件并分组
func readFileGroup(filePath string) (
	bool, []string) {
	var (
		flag   bool
		result []string
		dm     = make(map[string]int)
	)
	flag = utils.ReadFile(filePath, func(line string) bool {
		value := strings.TrimSpace(line)
		_, ok := dm[value]
		if !ok {
			dm[value] = 1
			return true
		} else {
			dm[value] = dm[value] + 1
			return true
		}
	})
	for k, v := range dm {
		result = append(result, fmt.Sprintf("%s : %d", k, v))
		fmt.Println(k, ":", v)
	}

	return flag, result
}
