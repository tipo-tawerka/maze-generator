package generator

import (
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
)

func TestDFSGeneratorGenerate(t *testing.T) {
	t.Parallel()
	generator := &dfsGenerator{}
	m := maze.NewMaze(5, 5)

	err := generator.Generate(&m)
	if err != nil {
		t.Errorf("DFS generation failed: %v", err)
	}

	if m.Rows() != 5 || m.Cols() != 5 {
		t.Error("Maze dimensions changed during generation")
	}
}

func TestDFSGeneratorSmallMaze(t *testing.T) {
	t.Parallel()
	generator := &dfsGenerator{}
	m := maze.NewMaze(3, 3)

	err := generator.Generate(&m)
	if err != nil {
		t.Errorf("DFS generation failed for small maze: %v", err)
	}
}

func TestDFSGeneratorLargeMaze(t *testing.T) {
	t.Parallel()
	generator := &dfsGenerator{}
	m := maze.NewMaze(21, 21)

	err := generator.Generate(&m)
	if err != nil {
		t.Errorf("DFS generation failed for large maze: %v", err)
	}
}
