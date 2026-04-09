package queue

import "github.com/tipo-tawerka/maze-generator/internal/domain/point"

type Queue struct {
	elements []point.Point
	start    int
}

func NewQueue() Queue {
	return Queue{
		elements: make([]point.Point, 0),
		start:    0,
	}
}

func (q *Queue) Add(p point.Point) {
	q.elements = append(q.elements, p)
}

func (q *Queue) Pop() point.Point {
	if q.start == len(q.elements) {
		panic("pop from empty queue")
	}
	p := q.elements[q.start]
	q.start++
	if q.start == len(q.elements) {
		clear(q.elements)
	}
	return p
}

func (q *Queue) IsEmpty() bool {
	return q.start >= len(q.elements)
}

func (q *Queue) Find(p point.Point) bool {
	for i := q.start; i < len(q.elements); i++ {
		if q.elements[i] == p {
			return true
		}
	}
	return false
}
