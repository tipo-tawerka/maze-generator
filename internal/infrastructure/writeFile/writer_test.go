package writeFile

import (
	"os"
	"testing"
)

func TestFileWriterWriteMaze(t *testing.T) {
	writer := FileWriter{}
	content := "Test maze content"

	tempFile, err := os.CreateTemp("", "test_output_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	_ = tempFile.Close()
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()

	err = writer.WriteMaze(content, tempFile.Name())
	if err != nil {
		t.Errorf("WriteMaze failed: %v", err)
	}

	writtenContent, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read written file: %v", err)
	}

	if string(writtenContent) != content {
		t.Errorf("Expected %q, got %q", content, string(writtenContent))
	}
}

func TestFileWriterWriteRealMaze(t *testing.T) {
	writer := FileWriter{}

	mazeContent := `#####
#S.F#
# ! #
# O #
#####`

	tempFile, err := os.CreateTemp("", "real_maze_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	_ = tempFile.Close()
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()

	err = writer.WriteMaze(mazeContent, tempFile.Name())
	if err != nil {
		t.Errorf("WriteMaze failed for real maze: %v", err)
	}

	writtenContent, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read written file: %v", err)
	}

	if string(writtenContent) != mazeContent {
		t.Errorf("Expected:\n%s\nGot:\n%s", mazeContent, string(writtenContent))
	}

	lines := []string{}
	currentLine := ""
	for _, char := range string(writtenContent) {
		if char == '\n' {
			lines = append(lines, currentLine)
			currentLine = ""
		} else {
			currentLine += string(char)
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	if len(lines) != 5 {
		t.Errorf("Expected 5 lines in maze, got %d", len(lines))
	}

	if lines[0] != "#####" {
		t.Errorf("First line should be '#####', got '%s'", lines[0])
	}
}

func TestFileWriterWriteMazeInvalidPath(t *testing.T) {
	writer := FileWriter{}
	content := "Test content"

	err := writer.WriteMaze(content, "/invalid/path/file.txt")
	if err == nil {
		t.Error("Should return error for invalid path")
	}
}

func TestFileWriterWriteMazeEmptyContent(t *testing.T) {
	writer := FileWriter{}

	tempFile, err := os.CreateTemp("", "empty_output_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	_ = tempFile.Close()
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()

	err = writer.WriteMaze("", tempFile.Name())
	if err != nil {
		t.Errorf("WriteMaze failed for empty content: %v", err)
	}

	writtenContent, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read written file: %v", err)
	}

	if string(writtenContent) != "" {
		t.Errorf("Expected empty content, got %q", string(writtenContent))
	}
}
