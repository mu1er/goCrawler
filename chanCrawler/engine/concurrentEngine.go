package engine

import (
	"log"

	"github.com/GoSpider/chanCrawler/model"
	"github.com/GoSpider/chanCrawler/scheduler"
	"github.com/GoSpider/crawler/fetcher"
)

type ConcurrentEngine struct {
	// 调度器
	Scheduler scheduler.Scheduler
	// 开启的 worker 数量
	WorkerCount int
}

func (e ConcurrentEngine) Run(seeds ...model.Request) {
	in := make(chan model.Request)
	out := make(chan model.ParserResult)
	//初始化调度器chan
	e.Scheduler.ConfigureMasterWorkerChan(in)
	// 创建workerCount个worker
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}
	// 将seeds的Request 添加到调度器chan
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out //阻塞获取
		for _, item := range result.Items {
			log.Printf("getItems,items:%v\n", item)
		}
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan model.Request, out chan model.ParserResult) {
	go func() {
		for {
			r := <-in // 阻塞等待获取
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r model.Request) (model.ParserResult, error) {
	log.Printf("fetching url:%s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch error, url: %s, err: %v", r.Url, err)
		return model.ParserResult{}, nil
	}
	return r.ParserFunc(body), nil
}
