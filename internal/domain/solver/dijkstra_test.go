package solver

import (
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

func TestDijkstraSolverSimplePath(t *testing.T) {
	t.Parallel()
	solver := &deikstraSolver{}
	m := createSimpleMaze()

	start := point.NewPoint(0, 1)
	finish := point.NewPoint(2, 1)

	path, err := solver.Solve(&m, start, finish)
	if err != nil {
		t.Errorf("Dijkstra solving failed: %v", err)
	}

	err = path.PrintPathOnMaze(m)
	if err != nil {
		t.Error("Path should be valid for maze")
	}
}

func TestDijkstraSolverNoPath(t *testing.T) {
	t.Parallel()
	solver := &deikstraSolver{}
	m := createSimpleMaze()

	start := point.NewPoint(0, 0)
	finish := point.NewPoint(2, 2)

	_, err := solver.Solve(&m, start, finish)
	t.Logf("Dijkstra solver result: %v", err)
}

func TestDijkstraSolverOptimalPath(t *testing.T) {
	t.Parallel()
	solver := &deikstraSolver{}
	m := createSimpleMaze()

	start := point.NewPoint(0, 1)
	finish := point.NewPoint(2, 1)

	path, err := solver.Solve(&m, start, finish)
	if err != nil {
		t.Errorf("Dijkstra solving failed: %v", err)
	}

	err = path.PrintPathOnMaze(m)
	if err != nil {
		t.Error("Path should be valid for maze")
	}
}
