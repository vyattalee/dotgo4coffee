package dispatcher2

import (
	"coffeeshop/worker2"
	"log"
)

// New returns a new dispatcher. A Dispatcher communicates between the client
// and the worker2. Its main job is to receive a job and share it on the WorkPool
// WorkPool is the link between the dispatcher and all the workers as
// the WorkPool of the dispatcher is common JobPool for all the workers
func New(pipelinenum int, workernum int) *disp {
	return &disp{
		Pipelines: make([]*Pipeline, pipelinenum),
		Workers:   make([]*worker2.Worker, workernum),
		JobChan:   make(worker2.JobChannel),
		Queue:     make(worker2.JobQueue),
	}
}

// disp is the link between the client and the workers
type disp struct {
	Pipelines []*Pipeline
	Workers   []*worker2.Worker // this is the list of workers that dispatcher tracks
	JobChan   worker2.JobChannel
	Queue     worker2.JobQueue // this is the shared JobPool between the workers
}

// Start creates pool of num count of workers.
func (d *disp) Start() *disp {

	pipeline1, pipeline2 := d.constructPipeline()

	d.Pipelines = append(d.Pipelines, pipeline1)
	d.Pipelines = append(d.Pipelines, pipeline2)

	select {
	case <-pipeline1.PipelineDone:
		log.Println("pipeline1-", pipeline1.Name, " done")
	}
	select {
	case <-pipeline2.PipelineDone:
		log.Println("pipeline2-", pipeline2.Name, " done")
	}

	l := len(d.Workers)
	for i := 1; i <= l; i++ {

		wrk := worker2.New(i, "default", make(worker2.JobChannel), d.Queue, make(chan struct{}))
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
		case job := <-d.JobChan: // listen to any submitted job on the WorkChan
			// wait for a worker2 to submit JobChan to Queue
			// note that this Queue is shared among all workers.
			// Whenever there is an available JobChan on Queue pull it
			JobChan := <-d.Queue

			// Once a jobChan is available, send the submitted Job on this JobChan
			JobChan <- job
		}
	}
}

func (d *disp) constructPipeline() (*Pipeline, *Pipeline) {
	pipeline1 := NewPipeline(1, "grindBean_espressoCoffee_pipeline")
	pipeline1.Machines <- &worker2.GrindBeanMachine{}
	pipeline1.Machines <- &worker2.EspressoCoffeeMachine{}
	close(pipeline1.Machines) //非常重要，不用了的channel务必关闭掉，否则就会有deadlock，继续等待channel接收数据

	pipeline2 := NewPipeline(2, "steamMilk_pipeline")
	pipeline2.Machines <- &worker2.SteamMilkMachine{}
	close(pipeline2.Machines) //非常重要，不用了的channel务必关闭掉，否则就会有deadlock，继续等待channel接收数据
	return pipeline1, pipeline2
}

//func (d *disp) SubmitJob(job worker2.Job) {
//	d.JobChan <- job
//}

//func (d *disp) submitPipeline(pipeline Pipeline) {
//	d.JobChan <- pipeline
//}
