package scheduler

import "github.com/GoSpider/chanCrawler/model"

type Scheduler interface {
	// 提交 Request 到调度器的 request 任务通道中
	Submit(request model.Request)
	// 初始化当前的调度器实例的 request 任务通道
	ConfigureMasterWorkerChan(chan model.Request)
}
