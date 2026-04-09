package minHeap

import "github.com/tipo-tawerka/maze-generator/internal/domain/point"

type elemHeap struct {
	point point.Point
	dist  int
}

type MinHeap struct {
	data []elemHeap
}

func NewMinHeap() MinHeap {
	return MinHeap{data: make([]elemHeap, 0, 32)}
}

func (mh *MinHeap) IsEmpty() bool {
	return len(mh.data) == 0
}

func (mh *MinHeap) Push(p point.Point, dist int) {
	mh.data = append(mh.data, elemHeap{point: p, dist: dist})
	mh.up(len(mh.data) - 1)
}

func (mh *MinHeap) Pop() point.Point {
	if mh.IsEmpty() {
		return point.Point{}
	}
	top := mh.data[0]
	lastIndex := len(mh.data) - 1
	mh.data[0] = mh.data[lastIndex]
	mh.data = mh.data[:lastIndex]
	mh.down(0)
	return top.point
}

func (mh *MinHeap) up(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2
	if mh.data[index].dist < mh.data[parent].dist {
		mh.data[index], mh.data[parent] = mh.data[parent], mh.data[index]
		mh.up(parent)
	}
}

func (mh *MinHeap) down(index int) {
	smallest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < len(mh.data) && mh.data[left].dist < mh.data[smallest].dist {
		smallest = left
	}
	if right < len(mh.data) && mh.data[right].dist < mh.data[smallest].dist {
		smallest = right
	}
	if smallest != index {
		mh.data[index], mh.data[smallest] = mh.data[smallest], mh.data[index]
		mh.down(smallest)
	}
}
