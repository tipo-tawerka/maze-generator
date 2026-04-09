package generator

import (
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
)

func TestPrimGeneratorGenerate(t *testing.T) {
	t.Parallel()
	generator := &primGenerator{}
	m := maze.NewMaze(5, 5)

	err := generator.Generate(&m)
	if err != nil {
		t.Errorf("Prim generation failed: %v", err)
	}

	if m.Rows() != 5 || m.Cols() != 5 {
		t.Error("Maze dimensions changed during generation")
	}
}

func TestPrimGeneratorSmallMaze(t *testing.T) {
	t.Parallel()
	generator := &primGenerator{}
	m := maze.NewMaze(3, 3)

	err := generator.Generate(&m)
	if err != nil {
		t.Errorf("Prim generation failed for small maze: %v", err)
	}
}

func TestPrimGeneratorMediumMaze(t *testing.T) {
	t.Parallel()
	generator := &primGenerator{}
	m := maze.NewMaze(11, 11)

	err := generator.Generate(&m)
	if err != nil {
		t.Errorf("Prim generation failed for medium maze: %v", err)
	}
}
