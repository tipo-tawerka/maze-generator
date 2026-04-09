package minHeap

import (
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

func TestNewMinHeap(t *testing.T) {
	t.Parallel()
	heap := NewMinHeap()

	if !heap.IsEmpty() {
		t.Error("New heap should be empty")
	}
}

func TestMinHeapPushPop(t *testing.T) {
	t.Parallel()
	heap := NewMinHeap()

	p1 := point.NewPoint(1, 1)
	p2 := point.NewPoint(2, 2)
	p3 := point.NewPoint(3, 3)

	heap.Push(p1, 10)
	heap.Push(p2, 5)
	heap.Push(p3, 15)

	if heap.IsEmpty() {
		t.Error("Heap should not be empty")
	}

	min1 := heap.Pop()
	if min1.X() != 2 || min1.Y() != 2 {
		t.Errorf("Expected point (2,2), got (%d,%d)", min1.X(), min1.Y())
	}
}

func TestMinHeapOrdering(t *testing.T) {
	t.Parallel()
	heap := NewMinHeap()

	distances := []int{30, 10, 20, 5, 40, 15}
	points := make([]point.Point, len(distances))

	for i, dist := range distances {
		p := point.NewPoint(i, i)
		points[i] = p
		heap.Push(p, dist)
	}

	first := heap.Pop()
	if first.X() != 3 || first.Y() != 3 {
		t.Errorf("First popped should be (3,3), got (%d,%d)", first.X(), first.Y())
	}

	second := heap.Pop()
	if second.X() != 1 || second.Y() != 1 {
		t.Errorf("Second popped should be (1,1), got (%d,%d)", second.X(), second.Y())
	}

	for !heap.IsEmpty() {
		_ = heap.Pop()
	}
}

func TestMinHeapEmptyOperations(t *testing.T) {
	t.Parallel()
	heap := NewMinHeap()

	if !heap.IsEmpty() {
		t.Error("Empty heap should return true for IsEmpty")
	}

	result := heap.Pop()
	if result.X() != 0 || result.Y() != 0 {
		t.Errorf("Pop on empty heap should return zero point, got (%d,%d)", result.X(), result.Y())
	}
}
