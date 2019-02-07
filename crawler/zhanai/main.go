package main

import (
	"github.com/GoSpider/crawler/engine"
	"github.com/GoSpider/crawler/zhanai/parser"
)

func main() {
	engine.Run(engine.Request{
		// 种子 Url
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
