package engine

// 请求任务封装
type Request struct {
	// 需要爬取的Url
	Url string
	// Url对应的解析函数
	ParserFunc func([]byte) ParserResult
}

// 解析结果
type ParserResult struct {
	// 解析出多个Request任务
	Requests []Request
	// 解析出来的实体
	Items []interface{}
}
