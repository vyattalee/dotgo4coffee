package main

import (
	"coffeeshop/dispatcher"
	"coffeeshop/worker"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	dd := dispatcher.New(16).Start()

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
			ID:        id,
			Name:      fmt.Sprintf("Job-::%s", name),
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
