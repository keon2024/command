package utils

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFile 读取文件，读取结果通过闭包方式传递，所以不体现在这里
// f中 返回值 bool表示这一行是否需要计数，比如有些方法里空串不计数
// ReadFile最外层返回表示整个读取的最终结果是否成功
func ReadFile(filePath string, f func(line string) bool) bool {
	var (
		cnt  int
		flag bool
	)
	defer func() {
		fmt.Println("总共提取", cnt, "条数据")
	}()
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件: %v\n", err)
		return flag
	}
	defer file.Close()

	// 创建一个新的 Reader 对象
	reader := bufio.NewReader(file)

	for {
		var line string
		// 读取一行
		line, err = reader.ReadString('\n')
		if err != nil {
			// 如果到达文件末尾或遇到错误，则退出循环
			if err.Error() != "EOF" {
				fmt.Println("读取文件时出错: %v", err)
				flag = false
			}
			break
		}
		if f(line) {
			// 成功才计数
			cnt++
		}
		flag = true

	}
	return flag
}
