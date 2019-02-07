package scheduler

import "github.com/GoSpider/chanCrawler/model"

type SimpleScheduler struct {
	workerChan chan model.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan model.Request) {
	s.workerChan = in
}

func (s *SimpleScheduler) Submit(request model.Request) {
	go func() { s.workerChan <- request }()
}
