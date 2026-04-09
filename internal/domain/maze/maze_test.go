package maze

import (
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

func TestNewMaze(t *testing.T) {
	t.Parallel()
	m := NewMaze(5, 7)
	if m.Rows() != 5 {
		t.Errorf("Expected 5 rows, got %d", m.Rows())
	}
	if m.Cols() != 7 {
		t.Errorf("Expected 7 cols, got %d", m.Cols())
	}
}

func TestMazeSetAndGetCell(t *testing.T) {
	t.Parallel()
	m := NewMaze(3, 3)

	wallCell, err := cellType.NewCellType(point.NewPoint(1, 1), cellType.Wall)
	if err != nil {
		t.Fatalf("Failed to create wall cell: %v", err)
	}

	m.SetCell(wallCell)

	retrievedCell := m.GetCell(point.NewPoint(1, 1))

	if retrievedCell.Print() != cellType.Wall {
		t.Errorf("Expected wall (%c), got %c", cellType.Wall, retrievedCell.Print())
	}
}

func TestMazeIsValid(t *testing.T) {
	t.Parallel()
	m := NewMaze(3, 5)

	testCases := []struct {
		x, y     int
		expected bool
	}{
		{0, 0, true},
		{4, 2, true},
		{2, 1, true},
		{-1, 0, false},
		{0, -1, false},
		{5, 0, false},
		{0, 3, false},
		{5, 3, false},
	}

	for _, tc := range testCases {
		result := m.IsValid(point.NewPoint(tc.x, tc.y))
		if result != tc.expected {
			t.Errorf("IsValid(%d, %d): expected %v, got %v",
				tc.x, tc.y, tc.expected, result)
		}
	}
}

func TestMazeGetFreeNeighbors(t *testing.T) {
	t.Parallel()
	m := NewMaze(5, 5)
	for i := 0; i < 5; i++ {
		topWall, _ := cellType.NewCellType(point.NewPoint(i, 0), cellType.Wall)
		bottomWall, _ := cellType.NewCellType(point.NewPoint(i, 4), cellType.Wall)
		m.SetCell(topWall)
		m.SetCell(bottomWall)
		leftWall, _ := cellType.NewCellType(point.NewPoint(0, i), cellType.Wall)
		rightWall, _ := cellType.NewCellType(point.NewPoint(4, i), cellType.Wall)
		m.SetCell(leftWall)
		m.SetCell(rightWall)
	}

	centerWall, _ := cellType.NewCellType(point.NewPoint(2, 2), cellType.Wall)
	m.SetCell(centerWall)

	for y := 1; y < 4; y++ {
		for x := 1; x < 4; x++ {
			if x == 2 && y == 2 {
				continue
			}
			emptyCell, _ := cellType.NewCellType(point.NewPoint(x, y), cellType.Empty)
			m.SetCell(emptyCell)
		}
	}

	neighbors := m.GetFreeNeighbors(point.NewPoint(1, 1))

	expectedNeighbors := 2
	if len(neighbors) != expectedNeighbors {
		t.Errorf("Expected %d neighbors for (1,1), got %d", expectedNeighbors, len(neighbors))
	}

	hasRight := false
	hasDown := false
	for _, neighbor := range neighbors {
		if neighbor.X() == 2 && neighbor.Y() == 1 {
			hasRight = true
		}
		if neighbor.X() == 1 && neighbor.Y() == 2 {
			hasDown = true
		}
	}

	if !hasRight {
		t.Error("Should have right neighbor (2,1)")
	}
	if !hasDown {
		t.Error("Should have down neighbor (1,2)")
	}
}

func TestMazeGetCellPanic(t *testing.T) {
	m := NewMaze(3, 3)

	defer func() {
		if r := recover(); r == nil {
			t.Error("GetCell should panic for invalid coordinates")
		}
	}()

	m.GetCell(point.NewPoint(5, 5))
}

func TestMazeSetCellPanic(t *testing.T) {
	m := NewMaze(3, 3)

	defer func() {
		if r := recover(); r == nil {
			t.Error("SetCell should panic for invalid coordinates")
		}
	}()

	invalidCell, _ := cellType.NewCellType(point.NewPoint(5, 5), cellType.Wall)
	m.SetCell(invalidCell)
}

func TestMazeComplexScenario(t *testing.T) {
	t.Parallel()
	m := NewMaze(4, 4)
	cellData := []struct {
		x, y     int
		cellType byte
	}{
		{0, 0, cellType.Start},
		{1, 0, cellType.Path},
		{2, 0, cellType.Highway},
		{3, 0, cellType.Finish},
		{0, 1, cellType.Wall},
		{1, 1, cellType.Wall},
		{2, 1, cellType.Wall},
		{3, 1, cellType.Pits},
		{0, 2, cellType.Path},
		{1, 2, cellType.Empty},
		{2, 2, cellType.Path},
		{3, 2, cellType.Wall},
		{0, 3, cellType.Wall},
		{1, 3, cellType.Path},
		{2, 3, cellType.Path},
		{3, 3, cellType.Wall},
	}

	for _, data := range cellData {
		cell, err := cellType.NewCellType(point.NewPoint(data.x, data.y), data.cellType)
		if err != nil {
			t.Fatalf("Failed to create cell at (%d,%d): %v", data.x, data.y, err)
		}
		m.SetCell(cell)
	}

	startCell := m.GetCell(point.NewPoint(0, 0))
	if startCell.Print() != cellType.Start {
		t.Errorf("Expected start cell, got %c", startCell.Print())
	}

	highwayCell := m.GetCell(point.NewPoint(2, 0))
	if highwayCell.Print() != cellType.Highway {
		t.Errorf("Expected highway cell, got %c", highwayCell.Print())
	}

	neighbors := m.GetFreeNeighbors(point.NewPoint(1, 2))
	expectedFreeNeighbors := 3
	if len(neighbors) != expectedFreeNeighbors {
		t.Errorf("Expected %d free neighbors for (1,2), got %d",
			expectedFreeNeighbors, len(neighbors))
	}
}
