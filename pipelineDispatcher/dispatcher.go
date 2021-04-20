package pipelineDispatcher

import (
	"coffeeshop/pipelineWorker"
)

// New returns a new dispatcher. A Dispatcher communicates between the client
// and the pipelineWorker. Its main job is to receive a job and share it on the WorkPool
// WorkPool is the link between the dispatcher and all the workers as
// the WorkPool of the dispatcher is common JobPool for all the workers
func New(num int) *disp {
	return &disp{
		Workers:  make([]*pipelineWorker.Worker, num),
		WorkChan: make(pipelineWorker.JobChannel),
		Queue:    make(pipelineWorker.JobQueue),
	}
}

// disp is the link between the client and the workers
type disp struct {
	Workers  []*pipelineWorker.Worker  // this is the list of workers that dispatcher tracks
	WorkChan pipelineWorker.JobChannel // client submits job to this channel
	Queue    pipelineWorker.JobQueue   // this is the shared JobPool between the workers
}

// Start creates pool of num count of workers.
func (d *disp) Start() *disp {
	l := len(d.Workers)
	for i := 1; i <= l; i++ {
		wrk := pipelineWorker.New(i, "pipeline-worker", make(pipelineWorker.JobChannel), d.Queue, make(chan struct{}))
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
		case job := <-d.WorkChan: // listen to any submitted job on the WorkChan
			// wait for a pipelineWorker to submit JobChan to JobQueue
			// note that this JobQueue is shared among all workers.
			// Whenever there is an available JobChan on JobQueue pull it
			jobChan := <-d.Queue

			// Once a jobChan is available, send the submitted Job on this JobChan
			jobChan <- job
		}
	}
}

func (d *disp) Submit(job pipelineWorker.Job) {
	d.WorkChan <- job
}
