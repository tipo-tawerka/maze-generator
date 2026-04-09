package readFile

import (
	"os"
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

func TestFileReaderReadMazeValidFile(t *testing.T) {
	reader := FileReader{}

	content := "###\n# #\n###\n"
	tempFile, err := os.CreateTemp("", "test_maze_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()

	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	_ = tempFile.Close()

	maze, err := reader.ReadMaze(tempFile.Name())
	if err != nil {
		t.Errorf("ReadMaze failed: %v", err)
	}

	if maze.Rows() != 3 || maze.Cols() != 3 {
		t.Errorf("Expected 3x3 maze, got %dx%d", maze.Rows(), maze.Cols())
	}
}

func TestFileReaderReadMazeWithContent(t *testing.T) {
	reader := FileReader{}
	content := "#####\n#S.F#\n# ! #\n# O #\n#####\n"
	tempFile, err := os.CreateTemp("", "content_maze_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()

	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	_ = tempFile.Close()

	maze, err := reader.ReadMaze(tempFile.Name())
	if err != nil {
		t.Errorf("ReadMaze failed: %v", err)
	}

	if maze.Rows() != 5 || maze.Cols() != 5 {
		t.Errorf("Expected 5x5 maze, got %dx%d", maze.Rows(), maze.Cols())
	}

	testCases := []struct {
		x, y     int
		expected byte
	}{
		{0, 0, cellType.Wall},
		{1, 1, cellType.Start},
		{2, 1, cellType.Path},
		{3, 1, cellType.Finish},
		{2, 2, cellType.Highway},
		{2, 3, cellType.Pits},
		{1, 2, cellType.Empty},
	}

	for _, tc := range testCases {
		cell := maze.GetCell(point.NewPoint(tc.x, tc.y))
		if cell.Print() != tc.expected {
			t.Errorf("Cell at (%d,%d): expected %c, got %c",
				tc.x, tc.y, tc.expected, cell.Print())
		}
	}
}

func TestFileReaderReadMazeNonexistentFile(t *testing.T) {
	reader := FileReader{}

	_, err := reader.ReadMaze("nonexistent_file.txt")
	if err == nil {
		t.Error("Should return error for nonexistent file")
	}
}

func TestFileReaderReadMazeEmptyFile(t *testing.T) {
	reader := FileReader{}

	tempFile, err := os.CreateTemp("", "empty_maze_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()
	_ = tempFile.Close()

	_, err = reader.ReadMaze(tempFile.Name())
	if err == nil {
		t.Error("Should return error for empty file")
	}
}
