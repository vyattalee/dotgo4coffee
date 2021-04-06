package worker1

import (
	"fmt"
	"log"
	"time"
)

type Machine interface {
	dojob(worker_id int, job Job) SemiFinishedProduct
	name() string
}

type SemiFinishedProduct struct {
	productId          int64
	productDescription string
}

var grindBeanTime, espressoCoffeeTime, steamMilkTime time.Duration = 1, 2, 3

type GrindBeanMachine struct {
}

func (g *GrindBeanMachine) dojob(worker_id int, job Job) SemiFinishedProduct {
	start := time.Now()
	prefix := fmt.Sprintf("Worker[%d]-Job[%d::%s]", worker_id, job.ID, job.Name)
	fmt.Println(prefix, "start to do job!")
	time.Sleep(time.Millisecond * grindBeanTime)
	end := time.Now()
	log.Print(end.Sub(start).Seconds())
	return SemiFinishedProduct{}
}

func (g *GrindBeanMachine) name() string {
	return "GrindBeanMachine"
}

type EspressoCoffeeMachine struct {
}

func (e *EspressoCoffeeMachine) dojob(worker_id int, job Job) SemiFinishedProduct {
	start := time.Now()
	prefix := fmt.Sprintf("Worker[%d]-Job[%d::%s]", worker_id, job.ID, job.Name)
	fmt.Println(prefix, "start to do job!")
	time.Sleep(time.Millisecond * espressoCoffeeTime)
	end := time.Now()
	log.Print(end.Sub(start).Seconds())
	return SemiFinishedProduct{}
}

func (e *EspressoCoffeeMachine) name() string {
	return "EspressoCoffeeMachine"
}

type SteamMilkMachine struct {
}

func (s *SteamMilkMachine) dojob(worker_id int, job Job) SemiFinishedProduct {
	start := time.Now()
	prefix := fmt.Sprintf("Worker[%d]-Job[%d::%s]", worker_id, job.ID, job.Name)
	fmt.Println(prefix, "start to do job!")
	time.Sleep(time.Millisecond * steamMilkTime)
	end := time.Now()
	log.Print(end.Sub(start).Seconds())
	return SemiFinishedProduct{}
}

func (s *SteamMilkMachine) name() string {
	return "SteamMilkMachine"
}
