package scheduler

import "go_pratice_code/learn_ccmouse_code/concurrent_crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// 会造成goroutine循环等待，导致goroutine死锁
	// s.workerChan <- r

	// 解决办法, 为每个request开1个goroutine，解决循环等待问题
	go func() { s.workerChan <- r }()
}

func (s *SimpleScheduler) workerReady(chan engine.Request) {
}

func (s *SimpleScheduler) workerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

// func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
// 	s.workerChan = c
// }
