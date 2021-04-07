//worker1 binding the unique machine, such as grindBean/espressoCoffee/steamMilk machine
//

package worker1

import (
	"log"
	"time"
)

// Job represents a single entity that should be processed.
// For example a struct that should be saved to database

type PipelineQueue chan chan Pipeline
type PipelineChannel chan Pipeline

// Worker is a a single processor. Typically its possible to
// start multiple workers for better throughput
type Worker struct {
	ID           int             // id of the worker
	PipelineChan PipelineChannel //a channel to machine list, a worker1 can deal with several machine in future
	Queue        PipelineQueue   // shared between all workers.
	Quit         chan struct{}   // a channel to quit working
}

func New(ID int, PipelineChan PipelineChannel, Queue PipelineQueue, Quit chan struct{}) *Worker {
	return &Worker{
		ID:           ID,
		PipelineChan: PipelineChan,
		Queue:        Queue,
		Quit:         Quit,
	}
}

func (wr *Worker) Start() {
	//c := &http.Client{Timeout: time.Millisecond * 15000}
	go func() {
		for {
			// when available, put the JobChan again on the JobPool
			// and wait to receive a job
			wr.Queue <- wr.PipelineChan
			select {
			case pipeline := <-wr.PipelineChan:
				// when a pipeline is received, process
				//callApi(job.ID, wr.ID, c)
				//callPipeline(pipeline)
				for machine := range pipeline.Machines {
					//l := len(pipeline.Machines)
					//for i := 1; i <= l; i++ {
					//for {
					//	select {
					//	case machine := <-pipeline.Machines:
					log.Println("pipeline-", pipeline.Name, " Machine-", machine.name(), " do job!")
					machine.dojob(wr.ID, Job{pipeline.ID<<6 + wr.ID, pipeline.Name + machine.name(), time.Now(), time.Now()})
					//default:
					//	continue
					//}
					//} //end for
				} //end for machine := range pipeline.Machines
				//dojob(wr.ID, job)
			case <-wr.Quit:
				// a signal on this channel means someone triggered
				// a shutdown for this worker
				close(wr.PipelineChan)
				return
			}
		}
	}()
}

// stop closes the Quit channel on the worker.
func (wr *Worker) Stop() {
	close(wr.Quit)
}

func callPipeline(pipeline Pipeline) {
	log.Println("pipeline:", pipeline.Name, "@", pipeline.ID, "  createAt:", pipeline.CreatedAt, "  updateAt:", pipeline.UpdatedAt)
}
