package worker

import (
	"fmt"
	"math/rand"
	"time"
)

type Jobable interface {
	dojob(worker_id int, job Job)
}

type SemiFinishedProduct struct {
	productId          int64
	productDescription string
}

var grindBeanTime, espressoCoffeeTime, steamMilkTime time.Duration = 1, 2, 3

type grindBeanJob struct {
	Job
}

func (g *grindBeanJob) dojob(worker_id int, job Job) SemiFinishedProduct {

	time.Sleep(time.Millisecond * grindBeanTime)
	return SemiFinishedProduct{}
}

type espressoCoffeeJob struct {
	Job
}

func (e *espressoCoffeeJob) dojob(worker_id int, job Job) SemiFinishedProduct {
	time.Sleep(time.Millisecond * espressoCoffeeTime)
	return SemiFinishedProduct{}
}

type steamMilkJob struct {
	Job
}

func (e *steamMilkJob) dojob(worker_id int, job Job) SemiFinishedProduct {
	time.Sleep(time.Millisecond * steamMilkTime)
	return SemiFinishedProduct{}
}

// Job represents a single entity that should be processed.
// For example a struct that should be saved to database
type Job struct {
	ID   int
	Name string
	//dojob	interface{}
	CreatedAt time.Time
	UpdatedAt time.Time
}

type JobChannel chan Job
type JobQueue chan chan Job

// Worker is a a single processor. Typically its possible to
// start multiple workers for better throughput
type Worker struct {
	ID      int           // id of the worker
	JobChan JobChannel    // a channel to receive single unit of work
	Queue   JobQueue      // shared between all workers.
	Quit    chan struct{} // a channel to quit working
}

func New(ID int, JobChan JobChannel, Queue JobQueue, Quit chan struct{}) *Worker {
	return &Worker{
		ID:      ID,
		JobChan: JobChan,
		Queue:   Queue,
		Quit:    Quit,
	}
}

func (wr *Worker) Start() {
	//c := &http.Client{Timeout: time.Millisecond * 15000}
	go func() {
		for {
			// when available, put the JobChan again on the JobPool
			// and wait to receive a job
			wr.Queue <- wr.JobChan
			select {
			case job := <-wr.JobChan:
				// when a job is received, process
				//callApi(job.ID, wr.ID, c)
				dojob(wr.ID, job)
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
}

func dojob(id int, job Job) {
	prefix := fmt.Sprintf("Worker[%d]-Job[%d::%s]", id, job.ID, job.Name)
	fmt.Println(prefix, "start to do job!")
	//time.Sleep(time.Millisecond * time.Duration(rand.Intn(100000)))
	time.Sleep(time.Second * time.Duration(rand.Intn(20)))

}

//
//func callApi(num, id int, c *http.Client) {
//	baseURL := "https://age-of-empires-2-api.herokuapp.com/api/v1/civilization/%d"
//
//	ur := fmt.Sprintf(baseURL, num)
//	req, err := http.NewRequest(http.MethodGet, ur, nil)
//	if err != nil {
//		//log.Printf("error creating a request for term %d :: error is %+v", num, err)
//		return
//	}
//	res, err := c.Do(req)
//	if err != nil {
//		//log.Printf("error querying for term %d :: error is %+v", num, err)
//		return
//	}
//	defer res.Body.Close()
//	_, err = ioutil.ReadAll(res.Body)
//	if err != nil {
//		//log.Printf("error reading response body :: error is %+v", err)
//		return
//	}
//	//log.Printf("%d  :: ok", id)
//}
