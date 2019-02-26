package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items []interface{}
}

// 处理nil parser情况
func NilParser([]byte) ParserResult{
	return ParserResult{}
}