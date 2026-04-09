// Package solver предоставляет различные алгоритмы для решения лабиринтов.
//
// Пакет содержит интерфейс Solver и его реализации для различных алгоритмов
// поиска кратчайшего пути в лабиринте, таких как алгоритм Дейкстры, A* и SPFA.
// Также включает фабричный метод для создания решателей по названию алгоритма
// и вспомогательные структуры для работы с данными вершин графа.
package solver

import (
	"fmt"

	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/path"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

// Solver определяет интерфейс для алгоритмов решения лабиринтов.
// Все реализации должны предоставлять метод Solve для поиска пути между двумя точками.
type Solver interface {
	// Solve находит кратчайший путь от начальной до конечной точки в лабиринте.
	// Возвращает найденный путь или ошибку, если путь не существует.
	Solve(maze *maze.Maze, start, finish point.Point) (path.Path, error)
}

// vertexData содержит информацию о вершине графа для алгоритмов поиска пути.
// Используется для хранения расстояния, родительской вершины и статуса посещения.
type vertexData struct {
	dist    int         // кратчайшее расстояние от начальной точки
	parent  point.Point // родительская вершина в кратчайшем пути
	visited bool        // флаг посещения вершины
}

// noParentPoint представляет отсутствие родительской вершины.
// Используется для обозначения начальной точки пути.
var noParentPoint = point.NewPoint(-1, -1)

// inf представляет бесконечное расстояние.
// Используется для инициализации расстояний в алгоритмах поиска пути.
var inf = 1 << 30

// Константы определяют доступные алгоритмы решения лабиринтов.
const (
	DijkstraAlgorithm = "dijkstra" // алгоритм Дейкстры для поиска кратчайшего пути
	AStarAlgorithm    = "astar"    // алгоритм A* с эвристикой для оптимизации поиска
	SPFAAlgorithm     = "spfa"     // алгоритм SPFA (Shortest Path Faster Algorithm)
)

// FabricateSolver создает решатель лабиринта для указанного алгоритма.
//
// Параметры:
//   - algorithm: название алгоритма решения ("dijkstra", "astar", "spfa")
//
// Возвращает:
//   - Solver: экземпляр решателя для указанного алгоритма
//   - error: ошибку, если алгоритм не поддерживается
//
// Поддерживаемые алгоритмы:
//   - "dijkstra" - алгоритм Дейкстры, гарантирует кратчайший путь
//   - "astar" - алгоритм A*, использует эвристику для ускорения поиска
//   - "spfa" - SPFA, эффективен для графов с отрицательными весами
func FabricateSolver(algorithm string) (Solver, error) {
	switch algorithm {
	case DijkstraAlgorithm:
		return &deikstraSolver{}, nil
	case AStarAlgorithm:
		return &aStarSolver{}, nil
	case SPFAAlgorithm:
		return &spfaSolver{}, nil
	default:
		return nil, fmt.Errorf("unknown algorithm: %s", algorithm)
	}
}
