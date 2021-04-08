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

	dd := dispatcher1.New(6).Start()

	//grindBean_espressoCoffee_machine := (worker1.grindBeanMachine)

	//terms := map[int]string{
	//	1: "steamMilk_pipeline",
	//	2: "grindBean_espressoCoffee_pipeline",
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
	//}

	//for id, name := range terms {
	//
	//	pipeline := worker1.NewPipeline(id, name)
	//	//pipeline.Machines[0] = worker1.Machine(&worker1.GrindBeanMachine{})
	//	//pipeline.Machines[1] = &worker1.EspressoCoffeeMachine{}
	//	pipeline.Machines <- &worker1.GrindBeanMachine{}
	//	pipeline.Machines <- &worker1.EspressoCoffeeMachine{}
	//
	//	dd.SubmitPipeline(*pipeline)
	//}

	pipeline1 := worker1.NewPipeline(1, "grindBean_espressoCoffee_pipeline")
	pipeline1.Machines <- &worker1.GrindBeanMachine{}
	pipeline1.Machines <- &worker1.EspressoCoffeeMachine{}

	pipeline2 := worker1.NewPipeline(2, "steamMilk_pipeline")
	pipeline2.Machines <- &worker1.SteamMilkMachine{}

	dd.SubmitPipeline(*pipeline1)
	dd.SubmitPipeline(*pipeline2)

	pipeline3 := worker1.NewPipeline(3, "3_pipeline")
	pipeline3.Machines <- &worker1.SteamMilkMachine{}
	pipeline4 := worker1.NewPipeline(4, "4_pipeline")
	pipeline4.Machines <- &worker1.SteamMilkMachine{}
	pipeline5 := worker1.NewPipeline(5, "5_pipeline")
	pipeline5.Machines <- &worker1.SteamMilkMachine{}
	pipeline6 := worker1.NewPipeline(6, "6_pipeline")
	pipeline6.Machines <- &worker1.SteamMilkMachine{}
	pipeline7 := worker1.NewPipeline(7, "7_pipeline")
	pipeline7.Machines <- &worker1.SteamMilkMachine{}
	pipeline8 := worker1.NewPipeline(8, "8_pipeline")
	pipeline8.Machines <- &worker1.SteamMilkMachine{}

	dd.SubmitPipeline(*pipeline3)
	dd.SubmitPipeline(*pipeline4)
	dd.SubmitPipeline(*pipeline5)
	dd.SubmitPipeline(*pipeline6)
	dd.SubmitPipeline(*pipeline7)
	//dd.SubmitPipeline(*pipeline8)

	end := time.Now()
	log.Print(end.Sub(start).Seconds())
}
