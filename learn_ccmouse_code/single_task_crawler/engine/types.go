package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// 处理nil parser情况
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
