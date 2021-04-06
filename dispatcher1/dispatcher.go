package dispatcher1

import (
	"coffeeshop/worker1"
)

// New returns a new dispatcher. A Dispatcher communicates between the client
// and the worker1. Its main job is to receive a job and share it on the WorkPool
// WorkPool is the link between the dispatcher and all the workers as
// the WorkPool of the dispatcher is common JobPool for all the workers
func New(num int) *disp {
	return &disp{
		Workers:      make([]*worker1.Worker, num),
		PipelineChan: make(worker1.PipelineChannel),
		Queue:        make(worker1.PipelineQueue),
	}
}

// disp is the link between the client and the workers
type disp struct {
	Workers      []*worker1.Worker // this is the list of workers that dispatcher tracks
	PipelineChan worker1.PipelineChannel
	Queue        worker1.PipelineQueue // this is the shared JobPool between the workers
}

// Start creates pool of num count of workers.
func (d *disp) Start() *disp {
	l := len(d.Workers)
	for i := 1; i <= l; i++ {
		wrk := worker1.New(i, make(worker1.PipelineChannel), d.Queue, make(chan struct{}))
		wrk.Start()
		d.Workers = append(d.Workers, wrk)
	}
	go d.process()
	return d
}

// process listens to a job submitted on WorkChan and
// relays it to the WorkPool. The WorkPool is shared between
// the workers.
func (d *disp) process() {
	for {
		select {
		case pipeline := <-d.PipelineChan: // listen to any submitted job on the WorkChan
			// wait for a worker1 to submit JobChan to Queue
			// note that this Queue is shared among all workers.
			// Whenever there is an available JobChan on Queue pull it
			pipelineChan := <-d.Queue

			// Once a jobChan is available, send the submitted Job on this JobChan
			pipelineChan <- pipeline
		}
	}
}

//func (d *disp) SubmitJob(job worker1.Job) {
//	d.WorkChan <- job
//}

func (d *disp) SubmitPipeline(pipeline worker1.Pipeline) {
	d.PipelineChan <- pipeline
}
