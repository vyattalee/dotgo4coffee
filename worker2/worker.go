//worker1 binding the unique machine, such as grindBean/espressoCoffee/steamMilk machine
//

package worker2

import (
	"log"
)

// Job represents a single entity that should be processed.
// For example a struct that should be saved to database

type JobQueue chan chan JobInterface

//type JobChannel chan Pipeline
type JobChannel chan JobInterface

// Worker is a a single processor. Typically its possible to
// start multiple workers for better throughput
type Worker struct {
	ID   int    // id of the worker
	Name string //name of the worker
	//PipelineChan JobChannel    //a channel to machine list, a worker1 can deal with several machine in future
	Queue   JobQueue      // shared between all workers.
	JobChan JobChannel    // a channel to
	Quit    chan struct{} // a channel to quit working
}

func New(ID int, Name string, JobChan JobChannel, Queue JobQueue, Quit chan struct{}) *Worker {
	return &Worker{
		ID:      ID,
		Name:    Name,
		JobChan: JobChan,
		Queue:   Queue,
		Quit:    Quit,
	}
}

func (wr *Worker) Start() {
	//c := &http.Client{Timeout: time.Millisecond * 15000}
	go func() {
		//defer wg.Done()
		for {

			// when available, put the JobChan again on the JobPool
			// and wait to receive a job
			wr.Queue <- wr.JobChan
			select {
			case job := <-wr.JobChan:
				// when a pipeline is received, process
				log.Println("Worker[", wr.ID, "::", wr.Name, "] do job[", job.JobID(), "::", job.JobName(), "]")
				//callApi(job.ID, wr.ID, c)
				//callPipeline(pipeline)
				//for machine := range pipeline.Machines {

				//l := len(pipeline.Machines)
				//for i := 1; i <= l; i++ {
				//for {
				//	select {
				//	case machine := <-pipeline.Machines:
				//log.Println("pipeline-", pipeline.Name, " Machine-", machine.name(), " do job!")
				//for job := range wr.JobChan{
				//	semiFinishedProduct := machine.dojob(*wr, job)
				//	log.Println("semiFinishedProduct[", semiFinishedProduct.productId, "::", semiFinishedProduct.productDescription, "]")
				//}
				//semiFinishedProduct := machine.dojob(*wr, Job{int64(pipeline.ID<<6 + wr.ID), pipeline.Name + machine.name(), time.Now(), time.Now()})
				//log.Println("semiFinishedProduct[", semiFinishedProduct.productId, "::", semiFinishedProduct.productDescription, "]")

				//default:
				//	continue
				//}
				//} //end for
				//} //end for machine := range pipeline.Machines
				//dojob(wr.ID, job)
				//wr.Stop()
				//pipeline.PipelineDone <- struct{}{}
				//close(pipeline.PipelineDone)
				wr.Quit <- struct{}{}
				wr.Stop()

			case <-wr.Quit:
				// a signal on this channel means someone triggered
				// a shutdown for this worker
				close(wr.JobChan)
				return
			}
		}

	}()
}

// stop closes the Quit channel on the worker.
func (wr *Worker) Stop() {
	close(wr.Quit)
	close(wr.JobChan)
}

//func (wr *Worker) SubmitJob(job JobInterface) {
//	wr.JobChan <- job
//}

//func callPipeline(pipeline dispatcher2.Pipeline) {
//	log.Println("pipeline:", pipeline.Name, "@", pipeline.ID, "  createAt:", pipeline.CreatedAt, "  updateAt:", pipeline.UpdatedAt)
//}
