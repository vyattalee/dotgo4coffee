package main

import (
	"coffeeshop/dispatcher"
	"coffeeshop/dispatcher1"
	"coffeeshop/dispatcher2"
	"coffeeshop/worker"
	"coffeeshop/worker2"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Mammal interface {
	Say()
}

type Dog struct{}

type Cat struct{}

type Human struct{}

func (d Dog) Say() {
	fmt.Println("woof")
}

func (c Cat) Say() {
	fmt.Println("meow")
}

func (h Human) Say() {
	fmt.Println("speak")
}

func test4PolymorphicInheritanceBYInterface() {
	var m Mammal
	var mammalchan chan Mammal

	m = &Dog{}
	m.Say()
	m = &Cat{}
	m.Say()
	m = &Human{}
	m.Say()

	mammalchan = make(chan Mammal, 5)
	mammalchan <- Human{}
	mammalchan <- Dog{}
	mammalchan <- Cat{}
	close(mammalchan) //remember channel should be closed after useing up

	for m := range mammalchan {
		//log.Println(m.Say())
		m.Say()
	}
}

func main() {
	//dispatch_worker()

	//dispatch_worker1()

	dispatch_worker2()

	//test4PolymorphicInheritanceBYInterface()

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

	//var wg sync.WaitGroup

	/*dd :=*/
	dispatcher1.New(6).Start()

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
	//	dd.submitPipeline(*pipeline)
	//}

	//pipeline3 := worker1.NewPipeline(3, "3_pipeline")
	//pipeline3.Machines <- &worker1.SteamMilkMachine{}
	//pipeline4 := worker1.NewPipeline(4, "4_pipeline")
	//pipeline4.Machines <- &worker1.SteamMilkMachine{}
	//pipeline5 := worker1.NewPipeline(5, "5_pipeline")
	//pipeline5.Machines <- &worker1.SteamMilkMachine{}
	//pipeline6 := worker1.NewPipeline(6, "6_pipeline")
	//pipeline6.Machines <- &worker1.SteamMilkMachine{}
	//pipeline7 := worker1.NewPipeline(7, "7_pipeline")
	//pipeline7.Machines <- &worker1.SteamMilkMachine{}
	//pipeline8 := worker1.NewPipeline(8, "8_pipeline")
	//pipeline8.Machines <- &worker1.SteamMilkMachine{}

	//dd.submitPipeline(*pipeline3)
	//dd.submitPipeline(*pipeline4)
	//dd.submitPipeline(*pipeline5)
	//dd.submitPipeline(*pipeline6)
	//dd.submitPipeline(*pipeline7)
	//dd.submitPipeline(*pipeline8)

	//l := len(dd.Workers)
	//for i := 0; i < l; i++ {
	//	<- dd.Workers[i].Quit
	//}

	end := time.Now()
	log.Print(end.Sub(start).Seconds())
	//wg.Wait()
	//fmt.Scanln()
}

func dispatch_worker2() {
	start := time.Now()

	//var wg sync.WaitGroup

	/**/
	dd := dispatcher2.New(6).Start()

	pipeline1 := worker2.NewPipeline(1, "grindBean_espressoCoffee_pipeline")
	pipeline1.Machines <- &worker2.GrindBeanMachine{}
	pipeline1.Machines <- &worker2.EspressoCoffeeMachine{}
	close(pipeline1.Machines) //非常重要，不用了的channel务必关闭掉，否则就会有deadlock，继续等待channel接收数据

	pipeline2 := worker2.NewPipeline(2, "steamMilk_pipeline")
	pipeline2.Machines <- &worker2.SteamMilkMachine{}
	close(pipeline2.Machines) //非常重要，不用了的channel务必关闭掉，否则就会有deadlock，继续等待channel接收数据

	dd.SubmitPipeline(*pipeline1)
	dd.SubmitPipeline(*pipeline2)

	dd.SubmitJob(worker2.Job{
		ID:        1000,
		Name:      fmt.Sprintf("Job-::%s", "normal order"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	dd.SubmitJob(worker2.SemiFinishedProduct{
		ProductId:          1001,
		ProductDescription: fmt.Sprintf("Job-::%s", "semiFinishedProduct"),
	})

	end := time.Now()
	log.Print(end.Sub(start).Seconds())
	//wg.Wait()
	//fmt.Scanln()
}
