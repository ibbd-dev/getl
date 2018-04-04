package common

const (
	TypeCSV   = "csv"
	TypeText  = "text"
	TypeMysql = "mysql"
)

// 参数
type Params struct {
	Input   string // 输入文件
	Output  string // 输出文件
	Silent  bool   // 关闭警告
	Verbose bool   // 打印进度信息
}
