package queue

import "fmt"

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

func (q *Queue) Add(n int) {
	q.data = append(q.data, n)
}

func (q *Queue) Remove() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}