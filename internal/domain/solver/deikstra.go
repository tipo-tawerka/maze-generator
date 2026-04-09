package solver

import (
	"errors"
	"fmt"

	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/path"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
	"github.com/tipo-tawerka/maze-generator/internal/domain/solver/minHeap"
)

// deikstraSolver реализует алгоритм Дейкстры для поиска кратчайшего пути в лабиринте.
//
// Алгоритм Дейкстры является классическим алгоритмом поиска кратчайшего пути
// в графе с неотрицательными весами рёбер. Гарантирует нахождение оптимального
// пути от начальной до конечной точки, учитывая стоимость прохождения через
// различные типы ячеек лабиринта.
//
// Особенности алгоритма:
//   - Гарантирует нахождение кратчайшего пути по стоимости
//   - Работает с различными типами ячеек (обычные, скоростные дороги, ямы)
//   - Использует приоритетную очередь (min-heap) для оптимизации
//   - Временная сложность: O((V + E) log V), где V - вершины, E - рёбра
//   - Подходит для лабиринтов с разными стоимостями прохождения
type deikstraSolver struct {
	maze     *maze.Maze                  // указатель на лабиринт для решения
	heap     minHeap.MinHeap             // приоритетная очередь для выбора следующей вершины
	vertexes map[point.Point]*vertexData // информация о каждой вершине (расстояние, родитель, статус)
}

// Solve находит кратчайший путь в лабиринте с использованием алгоритма Дейкстры.
//
// Метод реализует классический алгоритм Дейкстры с использованием приоритетной
// очереди. Начинает с начальной точки и постепенно исследует соседние вершины,
// всегда выбирая вершину с минимальным расстоянием. Гарантирует нахождение
// оптимального пути с учётом стоимости прохождения различных типов ячеек.
//
// Алгоритм работы:
//  1. Инициализирует все вершины с бесконечным расстоянием
//  2. Устанавливает расстояние до начальной точки равным 0
//  3. Пока приоритетная очередь не пуста:
//     - Извлекает вершину с минимальным расстоянием
//     - Обновляет расстояния до соседей, если найден лучший путь
//     - Добавляет соседей в приоритетную очередь
//  4. Восстанавливает путь по родительским связям
func (ds *deikstraSolver) Solve(maze *maze.Maze, start, finish point.Point) (path.Path, error) {
	if maze == nil {
		return path.NewPath(), errors.New("maze is nil")
	}
	ds.init(maze)
	if !ds.maze.IsValid(start) {
		return path.Path{}, errors.New("invalid start point")
	}
	if !ds.maze.IsValid(finish) {
		return path.Path{}, errors.New("invalid finish point")
	}
	ds.heap.Push(start, 0)
	ds.vertexes[start] = &vertexData{dist: 0, visited: false, parent: noParentPoint}
	for !ds.heap.IsEmpty() {
		currentPoint := ds.heap.Pop()
		if ds.vertexes[currentPoint].visited {
			continue
		}
		ds.vertexes[currentPoint].visited = true
		if currentPoint == finish {
			break
		}
		neighbors := ds.maze.GetFreeNeighbors(currentPoint)
		for _, neighbor := range neighbors {
			if ds.vertexes[neighbor].visited {
				continue
			}
			neighbourCell := ds.maze.GetCell(neighbor)
			newDist := ds.vertexes[currentPoint].dist + neighbourCell.GetCost()
			if newDist < ds.vertexes[neighbor].dist {
				ds.vertexes[neighbor].dist = newDist
				ds.vertexes[neighbor].parent = currentPoint
				ds.heap.Push(neighbor, newDist)
			}
		}
	}
	if !ds.vertexes[finish].visited {
		return path.Path{}, fmt.Errorf("no path from (%d, %d) to (%d, %d)",
			start.X(), start.Y(), finish.X(), finish.Y())
	}
	result := path.NewPath()
	for cur := finish; cur != noParentPoint; cur = ds.vertexes[cur].parent {
		result.AddPoint(cur)
	}
	return result, nil
}

// init инициализирует решатель Дейкстры и подготавливает структуры данных.
func (ds *deikstraSolver) init(maze *maze.Maze) {
	ds.maze = maze
	ds.heap = minHeap.NewMinHeap()
	ds.vertexes = make(map[point.Point]*vertexData, maze.Cols()*maze.Rows())
	for i := range maze.Rows() {
		for j := range maze.Cols() {
			p := point.NewPoint(j, i)
			ds.vertexes[p] = &vertexData{dist: inf, visited: false, parent: noParentPoint}
		}
	}
}
