package path

import (
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

func TestNewPath(t *testing.T) {
	t.Parallel()
	p := NewPath()
	if len(p.points) != 0 {
		t.Errorf("New path should be empty, got %d points", len(p.points))
	}
}

func TestPathAddPoint(t *testing.T) {
	t.Parallel()
	p := NewPath()
	pt1 := point.NewPoint(1, 1)
	pt2 := point.NewPoint(2, 2)

	p.AddPoint(pt1)
	if len(p.points) != 1 {
		t.Errorf("Expected 1 point, got %d", len(p.points))
	}

	p.AddPoint(pt2)
	if len(p.points) != 2 {
		t.Errorf("Expected 2 points, got %d", len(p.points))
	}
}

func TestPathPrintOnMaze(t *testing.T) {
	t.Parallel()
	testMaze := maze.NewMaze(3, 3)

	p := NewPath()
	pt1 := point.NewPoint(1, 1)
	pt2 := point.NewPoint(1, 2)

	p.AddPoint(pt1)
	p.AddPoint(pt2)

	err := p.PrintPathOnMaze(testMaze)
	if err != nil {
		t.Errorf("PrintPathOnMaze failed: %v", err)
	}
}

func TestPathPrintOnMazeEmpty(t *testing.T) {
	t.Parallel()
	testMaze := maze.NewMaze(3, 3)

	p := NewPath()

	err := p.PrintPathOnMaze(testMaze)
	if err != nil {
		t.Errorf("PrintPathOnMaze with empty path should not fail: %v", err)
	}
}
