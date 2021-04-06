package main

import (
	"coffeeshop/dispatcher"
	"coffeeshop/dispatcher1"
	"coffeeshop/worker"
	"coffeeshop/worker1"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	//dispatch_worker()

	dispatch_worker1()
}

func dispatch_worker() {
	start := time.Now()
	dd := dispatcher.New(10).Start()

	terms := map[int]string{
		1:  "grindBean",
		2:  "espressoCoffee",
		3:  "steamMilk",
		4:  "grindBean",
		5:  "espressoCoffee",
		6:  "steamMilk",
		7:  "grindBean",
		8:  "espressoCoffee",
		9:  "steamMilk",
		10: "grindBean",
		11: "espressoCoffee",
		12: "steamMilk",
		13: "grindBean",
		14: "espressoCoffee",
		15: "steamMilk",
		17: "grindBean",
		18: "espressoCoffee",
		19: "steamMilk",
		16: "coffeeLatte"}

	for id, name := range terms {
		dd.Submit(worker.Job{
			ID:   id,
			Name: fmt.Sprintf("Job-::%s", name),
			Dojob: func(id int, job worker.Job) {
				start := time.Now()
				prefix := fmt.Sprintf("##Worker[%d]-Job[%d::%s]", id, job.ID, job.Name)
				log.Print(prefix, "start to do job!")
				time.Sleep(time.Second * time.Duration(rand.Intn(10)))
				end := time.Now()
				log.Print(prefix, "finish job total time: ", end.Sub(start).Seconds())
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	//dd.Submit(worker.grindBeanJob{
	//	ID:        id,
	//	name:      fmt.Sprintf("Job-::%s", name),
	//	CreatedAt: time.Now(),
	//	UpdatedAt: time.Now(),
	//})

	end := time.Now()
	log.Print(end.Sub(start).Seconds())
}

func dispatch_worker1() {
	start := time.Now()

	dd := dispatcher1.New(16).Start()

	//grindBean_espressoCoffee_machine := (worker1.grindBeanMachine)

	terms := map[int]string{
		2: "grindBean_espressoCoffee_pipeline",
		//2:  "grindBean_espressoCoffee_pipeline",
		//3:  "steamMilk_pipeline",
		//4:  "grindBean_espressoCoffee_pipeline",
		//5:  "grindBean_espressoCoffee_pipeline",
		//6:  "steamMilk_pipeline",
		//7:  "grindBean_espressoCoffee_pipeline",
		//8:  "grindBean_espressoCoffee_pipeline",
		//9:  "steamMilk_pipeline",
		//10: "grindBean_espressoCoffee_pipeline",
		//11: "grindBean_espressoCoffee_pipeline",
		//12: "steamMilk_pipeline",
		//13: "grindBean_espressoCoffee_pipeline",
		//14: "grindBean_espressoCoffee_pipeline",
		//15: "steamMilk_pipeline",
		//16: "steamMilk_pipeline",
		//17: "grindBean_espressoCoffee_pipeline",
		//18: "grindBean_espressoCoffee_pipeline",
		1: "steamMilk_pipeline"}

	for id, name := range terms {

		pipeline := worker1.NewPipeline(id, name)
		pipeline.Machines <- &worker1.GrindBeanMachine{}
		pipeline.Machines <- &worker1.EspressoCoffeeMachine{}

		dd.SubmitPipeline(*pipeline)
	}

	end := time.Now()
	log.Print(end.Sub(start).Seconds())
}
