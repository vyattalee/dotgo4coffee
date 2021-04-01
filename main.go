package main

import (
	"coffeeshop/dispatcher"
	"coffeeshop/worker"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
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
	//	Name:      fmt.Sprintf("Job-::%s", name),
	//	CreatedAt: time.Now(),
	//	UpdatedAt: time.Now(),
	//})

	end := time.Now()
	log.Print(end.Sub(start).Seconds())
}
