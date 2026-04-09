package solver

import (
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

func TestSPFASolverSimplePath(t *testing.T) {
	t.Parallel()
	solver := &spfaSolver{}
	m := createSimpleMaze()

	start := point.NewPoint(0, 1)
	finish := point.NewPoint(2, 1)

	path, err := solver.Solve(&m, start, finish)
	if err != nil {
		t.Errorf("SPFA solving failed: %v", err)
	}

	err = path.PrintPathOnMaze(m)
	if err != nil {
		t.Error("Path should be valid for maze")
	}
}

func TestSPFASolverNoPath(t *testing.T) {
	t.Parallel()
	solver := &spfaSolver{}
	m := createSimpleMaze()

	start := point.NewPoint(0, 0)
	finish := point.NewPoint(2, 2)

	_, err := solver.Solve(&m, start, finish)
	t.Logf("SPFA solver result: %v", err)
}

func TestSPFASolverComplexMaze(t *testing.T) {
	t.Parallel()
	solver := &spfaSolver{}
	m := maze.NewMaze(5, 5)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i%2 == 1 || j%2 == 1 {
				cell, _ := cellType.NewCellType(point.NewPoint(i, j), cellType.Empty)
				m.SetCell(cell)
			}
		}
	}

	start := point.NewPoint(0, 1)
	finish := point.NewPoint(4, 1)

	path, err := solver.Solve(&m, start, finish)
	if err != nil {
		t.Errorf("SPFA solving failed for complex maze: %v", err)
	}

	err = path.PrintPathOnMaze(m)
	if err != nil {
		t.Error("Path should be valid for maze")
	}
}
