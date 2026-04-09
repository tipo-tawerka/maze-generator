package solving

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/printChar"
	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/readFile"
	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/writeFile"
)

const testMaze = `#######
#     #
# ### #
#   # #
# #   #
#     #
#######`

const simpleMaze = `#####
#   #
# # #
#   #
#####`

func createTestMazeFile(t *testing.T) string {
	tmpDir := t.TempDir()
	mazeFile := filepath.Join(tmpDir, "test_maze.txt")

	err := os.WriteFile(mazeFile, []byte(testMaze), 0644)
	if err != nil {
		t.Fatalf("Failed to create test maze file: %v", err)
	}

	return mazeFile
}

func createSimpleMazeFile(t *testing.T) string {
	tmpDir := t.TempDir()
	mazeFile := filepath.Join(tmpDir, "simple_maze.txt")

	err := os.WriteFile(mazeFile, []byte(simpleMaze), 0644)
	if err != nil {
		t.Fatalf("Failed to create simple maze file: %v", err)
	}

	return mazeFile
}

func TestSolveMazeAStar(t *testing.T) {
	mazeFile := createTestMazeFile(t)
	outputFile := filepath.Join(t.TempDir(), "output.txt")

	solveMaze := SolveMaze{
		Reader:  &readFile.FileReader{},
		Printer: &printChar.CharPrinter{},
		Writer:  &writeFile.FileWriter{},
	}

	err := solveMaze.Solve("astar", mazeFile, outputFile, "1,1", "5,5")
	if err != nil {
		t.Errorf("A* solve failed: %v", err)
	}
}

func TestSolveMazeWithPath(t *testing.T) {
	mazeFile := createSimpleMazeFile(t)
	outputFile := filepath.Join(t.TempDir(), "output.txt")

	solveMaze := SolveMaze{
		Reader:  &readFile.FileReader{},
		Printer: &printChar.CharPrinter{},
		Writer:  &writeFile.FileWriter{},
	}
	err := solveMaze.Solve("astar", mazeFile, outputFile, "1,1", "3,1")
	if err != nil {
		t.Fatalf("A* solve failed: %v", err)
	}

	result, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	resultStr := string(result)
	if !strings.Contains(resultStr, "S") {
		t.Error("Result should contain start point S")
	}
	if !strings.Contains(resultStr, "F") {
		t.Error("Result should contain finish point F")
	}
	if !strings.Contains(resultStr, "#####") {
		t.Error("Result should contain walls")
	}
	if !strings.Contains(resultStr, ".") {
		t.Error("Result should contain path .")
	}
}

func TestSolveMazeSPFA(t *testing.T) {
	mazeFile := createTestMazeFile(t)
	outputFile := filepath.Join(t.TempDir(), "output.txt")

	solveMaze := SolveMaze{
		Reader:  &readFile.FileReader{},
		Printer: &printChar.CharPrinter{},
		Writer:  &writeFile.FileWriter{},
	}

	err := solveMaze.Solve("spfa", mazeFile, outputFile, "1,1", "5,5")
	if err != nil {
		t.Errorf("SPFA solve failed: %v", err)
	}
}

func TestSolveMazeDijkstra(t *testing.T) {
	mazeFile := createTestMazeFile(t)
	outputFile := filepath.Join(t.TempDir(), "output.txt")

	solveMaze := SolveMaze{
		Reader:  &readFile.FileReader{},
		Printer: &printChar.CharPrinter{},
		Writer:  &writeFile.FileWriter{},
	}

	err := solveMaze.Solve("dijkstra", mazeFile, outputFile, "1,1", "5,5")
	if err != nil {
		t.Errorf("Dijkstra solve failed: %v", err)
	}
}

func TestSolveMazeInvalidAlgorithm(t *testing.T) {
	mazeFile := createTestMazeFile(t)
	outputFile := filepath.Join(t.TempDir(), "output.txt")

	solveMaze := SolveMaze{
		Reader:  &readFile.FileReader{},
		Printer: &printChar.CharPrinter{},
		Writer:  &writeFile.FileWriter{},
	}

	err := solveMaze.Solve("invalid", mazeFile, outputFile, "1,1", "5,5")
	if err == nil {
		t.Error("Should return error for invalid algorithm")
	}
}

func TestSolveMazeInvalidCoordinates(t *testing.T) {
	mazeFile := createTestMazeFile(t)
	outputFile := filepath.Join(t.TempDir(), "output.txt")

	solveMaze := SolveMaze{
		Reader:  &readFile.FileReader{},
		Printer: &printChar.CharPrinter{},
		Writer:  &writeFile.FileWriter{},
	}

	err := solveMaze.Solve("astar", mazeFile, outputFile, "invalid", "5,5")
	if err == nil {
		t.Error("Should return error for invalid start coordinates")
	}

	err = solveMaze.Solve("astar", mazeFile, outputFile, "1,1", "invalid")
	if err == nil {
		t.Error("Should return error for invalid end coordinates")
	}
}

func TestSolveMazeNonExistentFile(t *testing.T) {
	outputFile := filepath.Join(t.TempDir(), "output.txt")

	solveMaze := SolveMaze{
		Reader:  &readFile.FileReader{},
		Printer: &printChar.CharPrinter{},
		Writer:  &writeFile.FileWriter{},
	}

	err := solveMaze.Solve("astar", "nonexistent.txt", outputFile, "1,1", "5,5")
	if err == nil {
		t.Error("Should return error for non-existent file")
	}
}
