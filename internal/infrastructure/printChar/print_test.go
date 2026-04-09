package printChar

import (
	"strings"
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

func TestCharPrinterPrintMaze(t *testing.T) {
	t.Parallel()
	printer := CharPrinter{}
	m := maze.NewMaze(3, 3)

	wallCell, _ := cellType.NewCellType(point.NewPoint(0, 0), cellType.Wall)
	emptyCell, _ := cellType.NewCellType(point.NewPoint(1, 1), cellType.Empty)

	m.SetCell(wallCell)
	m.SetCell(emptyCell)

	result, err := printer.PrintMaze(&m)
	if err != nil {
		t.Errorf("PrintMaze failed: %v", err)
	}

	if result == "" {
		t.Error("Result should not be empty")
	}

	lines := strings.Split(strings.TrimSpace(result), "\n")
	if len(lines) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(lines))
	}
}

func TestCharPrinterSimpleMaze(t *testing.T) {
	t.Parallel()
	printer := CharPrinter{}
	m := maze.NewMaze(3, 3)
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			var cellChar byte
			if x == 0 || x == 2 || y == 0 || y == 2 {
				cellChar = cellType.Wall
			} else {
				cellChar = cellType.Empty
			}
			cell, _ := cellType.NewCellType(point.NewPoint(x, y), cellChar)
			m.SetCell(cell)
		}
	}

	result, err := printer.PrintMaze(&m)
	if err != nil {
		t.Errorf("PrintMaze failed: %v", err)
	}

	expected := "###\n# #\n###"
	result = strings.TrimSpace(result)

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestCharPrinterWithSpecialCells(t *testing.T) {
	t.Parallel()
	printer := CharPrinter{}
	m := maze.NewMaze(3, 5)

	testCells := []struct {
		x, y     int
		cellType byte
	}{
		{0, 0, cellType.Start},
		{1, 0, cellType.Path},
		{2, 0, cellType.Highway},
		{3, 0, cellType.Pits},
		{4, 0, cellType.Finish},
		{0, 1, cellType.Wall},
		{1, 1, cellType.Wall},
		{2, 1, cellType.Wall},
		{3, 1, cellType.Wall},
		{4, 1, cellType.Wall},
		{0, 2, cellType.Path},
		{1, 2, cellType.Pits},
		{2, 2, cellType.Empty},
		{3, 2, cellType.Empty},
		{4, 2, cellType.Empty},
	}

	for _, tc := range testCells {
		cell, _ := cellType.NewCellType(point.NewPoint(tc.x, tc.y), tc.cellType)
		m.SetCell(cell)
	}

	result, err := printer.PrintMaze(&m)
	if err != nil {
		t.Errorf("PrintMaze failed: %v", err)
	}

	expected := "S.!OF\n#####\n.O"
	result = strings.TrimSpace(result)

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestCharPrinterEmptyMaze(t *testing.T) {
	t.Parallel()
	printer := CharPrinter{}
	m := maze.NewMaze(2, 2)

	result, err := printer.PrintMaze(&m)
	if err != nil {
		t.Errorf("PrintMaze failed for empty maze: %v", err)
	}

	if result == "" {
		t.Error("Result should not be empty even for empty maze")
	}
}

func TestCharPrinterSingleCell(t *testing.T) {
	t.Parallel()
	printer := CharPrinter{}
	m := maze.NewMaze(1, 1)

	result, err := printer.PrintMaze(&m)
	if err != nil {
		t.Errorf("PrintMaze failed for single cell: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(result), "\n")
	if len(lines) != 1 {
		t.Errorf("Expected 1 line for single cell, got %d", len(lines))
	}
}
