package main

import (
	"github.com/GoSpider/chanCrawler/engine"
	"github.com/GoSpider/chanCrawler/model"
	"github.com/GoSpider/chanCrawler/scheduler"
	"github.com/GoSpider/chanCrawler/zhanai/parser"
)

func main() {
	engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 1000,
	}.Run(model.Request{
		// 种子 Url
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
