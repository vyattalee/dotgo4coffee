package worker1

import (
	"time"
)

type Pipeline struct {
	ID   int
	Name string
	//Machines  []*Machine
	Machines  chan Machine
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPipeline(ID int, Name string) *Pipeline {
	return &Pipeline{
		ID:   ID,
		Name: Name,
		//Machines:  make([]*Machine, 2),
		Machines:  make(chan Machine, 2),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
