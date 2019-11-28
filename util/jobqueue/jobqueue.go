package jobqueue

import (
	"fmt"
	"strconv"

	"github.com/phungvandat/onemilion/domain"
	"github.com/phungvandat/onemilion/service"
)

// Job struct
type Job struct {
	Payload domain.Payload
}

// JobQueue chan
type JobQueue chan Job

// NewJobQueue func
func NewJobQueue(maxQueue int) JobQueue {
	return make(chan Job, maxQueue)
}

// Worker struct
type worker struct {
	ID         string
	WorkerPool chan chan Job
	Job        chan Job
	Done       chan bool
}

// NewWorker func
func newWorker(workerPool chan chan Job, id string) worker {
	return worker{
		ID:         id,
		WorkerPool: workerPool,
		Job:        make(chan Job),
		Done:       make(chan bool),
	}
}

// Start func
func (w worker) start(svc service.Service) {
	go func() {
		for {
			w.WorkerPool <- w.Job
			select {
			case job := <-w.Job:
				if err := svc.Test.Test(job.Payload); err != nil {
					fmt.Println("Job ", job.Payload.Num, " failed by error ", err)
				}
			case <-w.Done:
				return
			}
		}
	}()
}

// Stop func
func (w worker) Stop() {
	go func() {
		w.Done <- true
	}()
}

// Dispatcher struct
type Dispatcher struct {
	maxWorker  int
	WorkerPool chan chan Job
}

// NewDispatcher func
func NewDispatcher(maxWorker int) *Dispatcher {
	pool := make(chan chan Job, maxWorker)
	return &Dispatcher{
		maxWorker:  maxWorker,
		WorkerPool: pool,
	}
}

// Run func
func (d *Dispatcher) Run(svc service.Service, queue JobQueue) {
	for i := 1; i <= d.maxWorker; i++ {
		worker := newWorker(d.WorkerPool, "w"+strconv.FormatInt(int64(i), 10))
		worker.start(svc)
	}
	go d.dispatch(queue)
}

func (d *Dispatcher) dispatch(queue JobQueue) {
	fmt.Println("Dispatcher started...")
	for {
		select {
		case job := <-queue:
			go func(job Job) {
				jobChan := <-d.WorkerPool
				jobChan <- job
			}(job)
		}
	}
}
