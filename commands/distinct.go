package commands

import (
	"command/utils"
	"fmt"
	"strconv"
	"strings"
)

type Distinct struct {
	File string
	Num  int
}

func (d *Distinct) Init(args []string) {
	d.File = args[1]
	if len(args) > 2 {
		num, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("num is not a number")
			return
		}
		d.Num = num
	}
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

	flag, result = readFileDistinct(d.File, d.Num)

	return flag, result
}

func (d *Distinct) Desc() {
	var desc = `distinct <file> <num>
	file: 文件路径
	num: 读取每行的字符数，按读取每行前 num 个字符内容去重`
	fmt.Println(desc)
}

// readFileDistinct 读取文件并去重
func readFileDistinct(filePath string, num int) (
	bool, []string) {
	var (
		flag   bool
		result []string
		dm     = make(map[string]int)
	)
	flag = utils.ReadFile(filePath, func(line string) bool {
		value := strings.TrimSpace(line)
		runes := []rune(value)
		if num > 0 && len(runes) > num {
			value = string(runes[:num]) // 截取前 num 个字符
		}
		if _, ok := dm[value]; !ok && value != "" {
			dm[value] = 1
			result = append(result, value)
			//fmt.Println(value)
			return true
		} else {
			dm[value]++
		}
		return false
	})
	for key, value := range dm {
		fmt.Printf("key: %s, count: %d\n", key, value)
	}

	return flag, result
}
