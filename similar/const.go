package similar

type Algo int

const (
	AlgoLevenshtein Algo = iota
)

type Similar struct {
	Algo           Algo   // 相似性算法
	Source, Target []rune // 需要计算相似性的两个字符串
}
