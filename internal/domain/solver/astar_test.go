package solver

import (
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

func createSimpleMaze() maze.Maze {
	m := maze.NewMaze(3, 3)

	emptyCell1, _ := cellType.NewCellType(point.NewPoint(0, 1), cellType.Empty)
	emptyCell2, _ := cellType.NewCellType(point.NewPoint(1, 1), cellType.Empty)
	emptyCell3, _ := cellType.NewCellType(point.NewPoint(2, 1), cellType.Empty)

	m.SetCell(emptyCell1)
	m.SetCell(emptyCell2)
	m.SetCell(emptyCell3)

	return m
}

func TestAStarSolverSimplePath(t *testing.T) {
	t.Parallel()
	solver := &aStarSolver{}
	m := createSimpleMaze()

	start := point.NewPoint(0, 1)
	finish := point.NewPoint(2, 1)

	path, err := solver.Solve(&m, start, finish)
	if err != nil {
		t.Errorf("AStar solving failed: %v", err)
	}

	err = path.PrintPathOnMaze(m)
	if err != nil {
		t.Error("Path should be valid for maze")
	}
}

func TestAStarSolverNoPath(t *testing.T) {
	t.Parallel()
	solver := &aStarSolver{}
	m := createSimpleMaze()

	start := point.NewPoint(0, 0)
	end := point.NewPoint(2, 2)

	_, err := solver.Solve(&m, start, end)
	t.Logf("A* solver result: %v", err)
}

func TestAStarSolverSameStartFinish(t *testing.T) {
	t.Parallel()
	solver := &aStarSolver{}
	m := createSimpleMaze()

	start := point.NewPoint(1, 1)
	finish := point.NewPoint(1, 1)

	path, err := solver.Solve(&m, start, finish)
	if err != nil {
		t.Errorf("AStar solving failed for same start/finish: %v", err)
	}

	err = path.PrintPathOnMaze(m)
	if err != nil {
		t.Error("Path should be valid")
	}
}
