package solver

import (
	"errors"
	"fmt"

	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/path"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
	"github.com/tipo-tawerka/maze-generator/internal/domain/solver/queue"
)

// spfaSolver реализует алгоритм SPFA для поиска кратчайшего пути в лабиринте.
//
// SPFA (Shortest Path Faster Algorithm) - это оптимизация алгоритма Беллмана-Форда,
// которая использует очередь для хранения вершин, требующих обновления расстояния.
// Алгоритм эффективен для графов с неотрицательными весами и может быть быстрее
// алгоритма Дейкстры в некоторых случаях, особенно для разрежённых графов.
type spfaSolver struct {
	queue    queue.Queue                 // обычная очередь для вершин, требующих обновления
	maze     *maze.Maze                  // указатель на лабиринт для решения
	vertexes map[point.Point]*vertexData // информация о каждой вершине (расстояние, родитель, статус)
}

// Solve находит кратчайший путь в лабиринте с использованием алгоритма SPFA.
//
// Метод реализует алгоритм SPFA (Shortest Path Faster Algorithm), который является
// оптимизированной версией алгоритма Беллмана-Форда. Использует обычную очередь
// для хранения вершин, которые могут потенциально улучшить расстояния до своих
// соседей. Алгоритм продолжает работу до тех пор, пока не останется вершин для обновления.
//
// Алгоритм работы:
//  1. Инициализирует все вершины с бесконечным расстоянием
//  2. Устанавливает расстояние до стартовой точки равным 0 и добавляет её в очередь
//  3. Пока очередь не пуста:
//     - Извлекает вершину из очереди
//     - Для каждого соседа проверяет возможность улучшения расстояния
//     - Если найден лучший путь, обновляет расстояние и добавляет соседа в очередь
//  4. Восстанавливает оптимальный путь по родительским связям
func (spfa *spfaSolver) Solve(maze *maze.Maze, start, finish point.Point) (path.Path, error) {
	if maze == nil {
		return path.NewPath(), errors.New("maze is nil")
	}
	spfa.init(maze)
	spfa.vertexes[start] = &vertexData{dist: 0, visited: true, parent: noParentPoint}
	spfa.queue.Add(start)
	for !spfa.queue.IsEmpty() {
		currentPoint := spfa.queue.Pop()
		currentVertex := spfa.vertexes[currentPoint]
		neighbors := maze.GetFreeNeighbors(currentPoint)
		for _, neighbor := range neighbors {
			neighbourCell := spfa.maze.GetCell(neighbor)
			weight := neighbourCell.GetCost()
			if currentVertex.dist+weight < spfa.vertexes[neighbor].dist {
				spfa.vertexes[neighbor].dist = currentVertex.dist + weight
				spfa.vertexes[neighbor].parent = currentPoint
				if !spfa.queue.Find(neighbor) {
					spfa.queue.Add(neighbor)
					spfa.vertexes[neighbor].visited = true
				}
			}
		}
	}
	if !spfa.vertexes[finish].visited {
		return path.Path{}, fmt.Errorf("no path from (%d, %d) to (%d, %d)",
			start.X(), start.Y(), finish.X(), finish.Y())
	}
	resPath := path.NewPath()
	for cur := finish; cur != noParentPoint; cur = spfa.vertexes[cur].parent {
		resPath.AddPoint(cur)
	}
	return resPath, nil
}

// init инициализирует решатель SPFA и подготавливает структуры данных.
func (spfa *spfaSolver) init(maze *maze.Maze) {
	spfa.maze = maze
	spfa.queue = queue.NewQueue()
	spfa.vertexes = make(map[point.Point]*vertexData, maze.Cols()*maze.Rows())
	inf := 1 << 30
	for i := range maze.Rows() {
		for j := range maze.Cols() {
			p := point.NewPoint(j, i)
			spfa.vertexes[p] = &vertexData{dist: inf, visited: false, parent: noParentPoint}
		}
	}
}
