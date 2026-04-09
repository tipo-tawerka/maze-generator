package generating

import (
	"os"
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/printChar"
	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/writeFile"
)

func TestGenerateMazeDFS(t *testing.T) {
	generator := GenerateMaze{
		Printer: &printChar.CharPrinter{},
		Writer:  &writeFile.FileWriter{},
	}

	tempFile, err := os.CreateTemp("", "test_generate_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	_ = tempFile.Close()
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()

	err = generator.Generate("dfs", tempFile.Name(), "5", "5")
	if err != nil {
		t.Errorf("Generate failed: %v", err)
	}

	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	if len(content) == 0 {
		t.Error("Generated maze should not be empty")
	}
}

func TestGenerateMazePrim(t *testing.T) {
	generator := GenerateMaze{
		Printer: &printChar.CharPrinter{},
		Writer:  &writeFile.FileWriter{},
	}

	tempFile, err := os.CreateTemp("", "test_prim_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	_ = tempFile.Close()
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()

	err = generator.Generate("prim", tempFile.Name(), "7", "7")
	if err != nil {
		t.Errorf("Prim generate failed: %v", err)
	}
}

func TestGenerateMazeInvalidAlgorithm(t *testing.T) {
	generator := GenerateMaze{
		Printer: &printChar.CharPrinter{},
		Writer:  &writeFile.FileWriter{},
	}

	err := generator.Generate("invalid", "output.txt", "5", "5")
	if err == nil {
		t.Error("Should return error for invalid algorithm")
	}
}

func TestGenerateMazeInvalidDimensions(t *testing.T) {
	generator := GenerateMaze{
		Printer: &printChar.CharPrinter{},
		Writer:  &writeFile.FileWriter{},
	}

	err := generator.Generate("dfs", "output.txt", "invalid", "5")
	if err == nil {
		t.Error("Should return error for invalid width")
	}

	err = generator.Generate("dfs", "output.txt", "5", "invalid")
	if err == nil {
		t.Error("Should return error for invalid height")
	}
}
