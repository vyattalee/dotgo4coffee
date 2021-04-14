package dispatcher2

import (
	"coffeeshop/worker2"
	"time"
)

type Pipeline struct {
	ID   int
	Name string
	//Machines  []*Machine
	Machines                chan worker2.Machine
	PipelineDone            chan struct{}
	SemiFinishedProductList chan worker2.SemiFinishedProduct
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

func NewPipeline(ID int, Name string) *Pipeline {
	return &Pipeline{
		ID:   ID,
		Name: Name,
		//Machines:  make([]*Machine, 2),
		Machines:                make(chan worker2.Machine, 2),
		PipelineDone:            make(chan struct{}),
		SemiFinishedProductList: make(chan worker2.SemiFinishedProduct, 10),
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}
}
