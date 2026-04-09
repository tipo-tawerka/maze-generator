package cellType

import (
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

func TestNewCellTypeEmpty(t *testing.T) {
	t.Parallel()
	p := point.NewPoint(1, 2)
	cell, err := NewCellType(p, Empty)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !cell.IsEmpty() {
		t.Error("Cell should be empty")
	}
	if cell.IsWall() {
		t.Error("Cell should not be wall")
	}
}

func TestNewCellTypeWall(t *testing.T) {
	t.Parallel()
	p := point.NewPoint(0, 0)
	cell, err := NewCellType(p, Wall)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if cell.IsEmpty() {
		t.Error("Cell should not be empty")
	}
	if !cell.IsWall() {
		t.Error("Cell should be wall")
	}
}

func TestCellTypeGetPoint(t *testing.T) {
	t.Parallel()
	p := point.NewPoint(5, 10)
	cell, _ := NewCellType(p, Empty)
	cellPoint := cell.Point
	if cellPoint.X() != 5 || cellPoint.Y() != 10 {
		t.Errorf("Expected point (5,10), got (%d,%d)", cellPoint.X(), cellPoint.Y())
	}
}

func TestCellTypePrint(t *testing.T) {
	t.Parallel()
	p := point.NewPoint(0, 0)

	emptyCell, _ := NewCellType(p, Empty)
	if emptyCell.Print() != ' ' {
		t.Errorf("Empty cell should print space, got %c", emptyCell.Print())
	}

	wallCell, _ := NewCellType(p, Wall)
	if wallCell.Print() != '#' {
		t.Errorf("Wall cell should print #, got %c", wallCell.Print())
	}
}

func TestCellTypeGetCost(t *testing.T) {
	t.Parallel()
	p := point.NewPoint(0, 0)
	cell, _ := NewCellType(p, Empty)
	if cell.GetCost() != 2 {
		t.Errorf("Expected cost 2, got %d", cell.GetCost())
	}
}
