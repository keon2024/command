package define

type Command interface {
	// Init 初始化
	Init(args []string)
	// Check 前置检查
	Check(args []string) bool
	// Exec 命令执行
	Exec() (bool, []string)
	// Desc 命令描述
	Desc()
}
