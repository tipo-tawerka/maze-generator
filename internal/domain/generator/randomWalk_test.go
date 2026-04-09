package generator

import (
	"testing"

	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
)

func TestRandomWalkGeneratorGenerate(t *testing.T) {
	t.Parallel()
	generator := &randomWalkGenerator{}
	m := maze.NewMaze(7, 7)

	err := generator.Generate(&m)
	if err != nil {
		t.Errorf("RandomWalk generation failed: %v", err)
	}

	if m.Rows() != 7 || m.Cols() != 7 {
		t.Error("Maze dimensions changed during generation")
	}
}

func TestRandomWalkGeneratorSmallMaze(t *testing.T) {
	t.Parallel()
	generator := &randomWalkGenerator{}
	m := maze.NewMaze(5, 5)

	err := generator.Generate(&m)
	if err != nil {
		t.Errorf("RandomWalk generation failed for small maze: %v", err)
	}
}

func TestRandomWalkGeneratorConsistency(t *testing.T) {
	t.Parallel()
	generator1 := &randomWalkGenerator{}
	generator2 := &randomWalkGenerator{}
	m1 := maze.NewMaze(9, 9)
	m2 := maze.NewMaze(9, 9)

	err1 := generator1.Generate(&m1)
	err2 := generator2.Generate(&m2)

	if err1 != nil || err2 != nil {
		t.Error("RandomWalk generation should not fail")
	}
}
