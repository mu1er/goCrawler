package engine

import (
	"fmt"
	"log"

	"github.com/GoSpider/crawler/fetcher"
)

func Run(seed ...Request) {
	// Request 队列
	var requests []Request
	// 将seed Request 放到Requst 初始化
	for _, r := range seed {
		requests = append(requests, r)
	}
	// 执行任务
	for len(requests) > 0 {
		// 获取地一个request 并且从 requests移出， 实现一个队列功能
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
		//error log
		if err != nil {
			log.Printf("fetch error, url: %s, err: %v", r.Url, err)
			continue
		}
		// 利用parserResult 对获取到的内容进行解析
		parserResult := r.ParserFunc(body)
		// 将解析到的request 添加到requests
		requests = append(requests, parserResult.Requests...)
		// 遍历 打印解析出来的实体

		for _, item := range parserResult.Items {
			fmt.Printf("getItems, url: %s, items: %v", r.Url, item)
		}
	}
}
