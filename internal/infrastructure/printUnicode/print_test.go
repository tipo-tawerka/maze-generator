package printUnicode

import (
	"strings"
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

func TestUnicodePrinterPrintMaze(t *testing.T) {
	printer := UnicodePrinter{}
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
	if len(lines) < 1 {
		t.Errorf("Expected at least 1 line, got %d", len(lines))
	}
}

func TestUnicodePrinterSimpleMaze(t *testing.T) {
	printer := UnicodePrinter{}
	m := maze.NewMaze(5, 5)

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			var cellChar byte
			if x == 0 || x == 4 || y == 0 || y == 4 {
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

	if !strings.Contains(result, "┌") && !strings.Contains(result, "─") && !strings.Contains(result, "┐") {
		t.Error("Result should contain Unicode box drawing characters")
	}

	lines := strings.Split(strings.TrimSpace(result), "\n")
	if len(lines) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(lines))
	}
}

func TestUnicodePrinterWithConnections(t *testing.T) {
	printer := UnicodePrinter{}
	m := maze.NewMaze(3, 3)

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			cell, _ := cellType.NewCellType(point.NewPoint(x, y), cellType.Wall)
			m.SetCell(cell)
		}
	}

	result, err := printer.PrintMaze(&m)
	if err != nil {
		t.Errorf("PrintMaze failed: %v", err)
	}

	hasConnectors := strings.Contains(result, "┬") ||
		strings.Contains(result, "┴") ||
		strings.Contains(result, "├") ||
		strings.Contains(result, "┤") ||
		strings.Contains(result, "┼") ||
		strings.Contains(result, "+")

	if !hasConnectors {
		t.Error("Result should contain connection Unicode characters for all-wall maze")
	}
}

func TestUnicodePrinterEmptySpaces(t *testing.T) {
	printer := UnicodePrinter{}
	m := maze.NewMaze(5, 5)

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			cell, _ := cellType.NewCellType(point.NewPoint(x, y), cellType.Empty)
			m.SetCell(cell)
		}
	}

	result, err := printer.PrintMaze(&m)
	if err != nil {
		t.Errorf("PrintMaze failed: %v", err)
	}

	if result == "" {
		t.Error("Result should not be completely empty")
	}

	hasSpacesOrNewlines := strings.Count(result, " ") > 0 || strings.Count(result, "\n") > 0
	if !hasSpacesOrNewlines {
		t.Error("Empty maze should contain spaces or newlines")
	}
}

func TestUnicodePrinterSingleWall(t *testing.T) {
	printer := UnicodePrinter{}
	m := maze.NewMaze(3, 3)

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			var cellChar byte
			if x == 1 && y == 1 {
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

	if result == "" {
		t.Error("Result should not be empty for single wall maze")
	}
}

func TestUnicodePrinterEdgeCases(t *testing.T) {
	printer := UnicodePrinter{}

	m1 := maze.NewMaze(1, 1)
	cell, _ := cellType.NewCellType(point.NewPoint(0, 0), cellType.Wall)
	m1.SetCell(cell)

	result1, err := printer.PrintMaze(&m1)
	if err != nil {
		t.Errorf("PrintMaze failed for 1x1 maze: %v", err)
	}
	if result1 == "" {
		t.Error("Result should not be empty for 1x1 maze")
	}

	m2 := maze.NewMaze(2, 2)
	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			cell, _ := cellType.NewCellType(point.NewPoint(x, y), cellType.Wall)
			m2.SetCell(cell)
		}
	}

	result2, err := printer.PrintMaze(&m2)
	if err != nil {
		t.Errorf("PrintMaze failed for 2x2 maze: %v", err)
	}
	if result2 == "" {
		t.Error("Result should not be empty for 2x2 maze")
	}
}
