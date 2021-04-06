package worker1

import "time"

type Job struct {
	ID        int
	Name      string
	Dojob     func(id int, job Job)
	CreatedAt time.Time
	UpdatedAt time.Time
}
