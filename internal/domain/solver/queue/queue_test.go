package queue

import (
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

func TestNewQueue(t *testing.T) {
	t.Parallel()
	q := NewQueue()

	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
}

func TestQueueAddPop(t *testing.T) {
	t.Parallel()
	q := NewQueue()

	p1 := point.NewPoint(1, 1)
	p2 := point.NewPoint(2, 2)
	p3 := point.NewPoint(3, 3)

	q.Add(p1)
	q.Add(p2)
	q.Add(p3)

	if q.IsEmpty() {
		t.Error("Queue should not be empty after adding elements")
	}

	first := q.Pop()
	if first.X() != 1 || first.Y() != 1 {
		t.Errorf("Expected first element (1,1), got (%d,%d)", first.X(), first.Y())
	}

	second := q.Pop()
	if second.X() != 2 || second.Y() != 2 {
		t.Errorf("Expected second element (2,2), got (%d,%d)", second.X(), second.Y())
	}

	third := q.Pop()
	if third.X() != 3 || third.Y() != 3 {
		t.Errorf("Expected third element (3,3), got (%d,%d)", third.X(), third.Y())
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty after popping all elements")
	}
}

func TestQueueFind(t *testing.T) {
	t.Parallel()
	q := NewQueue()

	p1 := point.NewPoint(1, 1)
	p2 := point.NewPoint(2, 2)
	p3 := point.NewPoint(3, 3)
	p4 := point.NewPoint(4, 4)

	q.Add(p1)
	q.Add(p2)
	q.Add(p3)

	if !q.Find(p1) {
		t.Error("Should find p1 in queue")
	}

	if !q.Find(p2) {
		t.Error("Should find p2 in queue")
	}

	if !q.Find(p3) {
		t.Error("Should find p3 in queue")
	}

	if q.Find(p4) {
		t.Error("Should not find p4 in queue")
	}
}

func TestQueueFindAfterPop(t *testing.T) {
	t.Parallel()
	q := NewQueue()

	p1 := point.NewPoint(1, 1)
	p2 := point.NewPoint(2, 2)
	p3 := point.NewPoint(3, 3)

	q.Add(p1)
	q.Add(p2)
	q.Add(p3)

	popped := q.Pop()
	if popped.X() != 1 || popped.Y() != 1 {
		t.Errorf("Expected popped element (1,1), got (%d,%d)", popped.X(), popped.Y())
	}

	if q.Find(p1) {
		t.Error("Should not find p1 after it was popped")
	}

	if !q.Find(p2) {
		t.Error("Should still find p2 in queue")
	}

	if !q.Find(p3) {
		t.Error("Should still find p3 in queue")
	}
}

func TestQueueMultiplePopFind(t *testing.T) {
	t.Parallel()
	q := NewQueue()

	for i := 0; i < 5; i++ {
		p := point.NewPoint(i, i)
		q.Add(p)
	}

	_ = q.Pop()
	_ = q.Pop()

	if q.Find(point.NewPoint(0, 0)) {
		t.Error("Should not find (0,0) after it was popped")
	}

	if q.Find(point.NewPoint(1, 1)) {
		t.Error("Should not find (1,1) after it was popped")
	}

	if !q.Find(point.NewPoint(2, 2)) {
		t.Error("Should find (2,2) in queue")
	}

	if !q.Find(point.NewPoint(4, 4)) {
		t.Error("Should find (4,4) in queue")
	}
}

func TestQueueEmptyState(t *testing.T) {
	t.Parallel()
	q := NewQueue()

	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	p := point.NewPoint(1, 1)
	if q.Find(p) {
		t.Error("Should not find element in empty queue")
	}

	q.Add(p)
	if q.IsEmpty() {
		t.Error("Queue should not be empty after adding element")
	}

	_ = q.Pop()
	if !q.IsEmpty() {
		t.Error("Queue should be empty after popping the only element")
	}
}
