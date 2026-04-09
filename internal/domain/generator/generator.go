// Package generator предоставляет различные алгоритмы генерации лабиринтов.
//
// Пакет содержит интерфейс Generator и его реализации для различных алгоритмов
// генерации лабиринтов, таких как поиск в глубину (DFS), алгоритм Прима
// и случайное блуждание. Также включает фабричный метод для создания
// генераторов по названию алгоритма.
package generator

import (
	"fmt"

	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
)

// Generator определяет интерфейс для алгоритмов генерации лабиринтов.
// Все реализации должны предоставлять метод Generate для создания лабиринта.
type Generator interface {
	// Generate генерирует лабиринт, изменяя переданную структуру maze.
	// Возвращает ошибку в случае проблем с генерацией.
	Generate(maze *maze.Maze) error
}

// Константы определяют доступные алгоритмы генерации лабиринтов.
const (
	DfsAlgorithm      = "dfs"      // поиск в глубину (Depth-First Search)
	PrimAlgorithm     = "prim"     // алгоритм Прима для создания минимального остовного дерева
	RandWalkAlgorithm = "randwalk" // случайное блуждание (Random Walk)
)

// FabricateGenerator создает генератор лабиринта для указанного алгоритма.
//
// Параметры:
//   - algorithm: название алгоритма генерации ("dfs", "prim", "randwalk")
//
// Возвращает:
//   - Generator: экземпляр генератора для указанного алгоритма
//   - error: ошибку, если алгоритм не поддерживается
//
// Поддерживаемые алгоритмы:
//   - "dfs" - поиск в глубину, создает лабиринты с длинными извилистыми путями
//   - "prim" - алгоритм Прима, создает более равномерные лабиринты
//   - "randwalk" - случайное блуждание, создает лабиринты со случайной структурой
func FabricateGenerator(algorithm string) (Generator, error) {
	switch algorithm {
	case PrimAlgorithm:
		return &primGenerator{}, nil
	case DfsAlgorithm:
		return &dfsGenerator{}, nil
	case RandWalkAlgorithm:
		return &randomWalkGenerator{}, nil
	default:
		return nil, fmt.Errorf("unknown generator algorithm: %s", algorithm)
	}
}

// direction представляет направление движения в лабиринте.
// Используется для определения соседних ячеек при генерации.
type direction struct {
	x, y int // смещение по координатам x и y
}
