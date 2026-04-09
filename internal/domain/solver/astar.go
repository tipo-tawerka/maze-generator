package solver

import (
	"errors"
	"fmt"

	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/path"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
	"github.com/tipo-tawerka/maze-generator/internal/domain/solver/minHeap"
)

// aStarSolver реализует алгоритм A* для поиска кратчайшего пути в лабиринте.
//
// Алгоритм A* (A-star) является эвристическим алгоритмом поиска кратчайшего пути,
// который использует функцию оценки f(n) = g(n) + h(n), где g(n) - фактическое
// расстояние от начала до текущей вершины, а h(n) - эвристическая оценка
// расстояния от текущей вершины до цели. Использует манхэттенское расстояние
// как эвристику, что делает его более эффективным чем алгоритм Дейкстры.
type aStarSolver struct {
	maze     *maze.Maze                  // указатель на лабиринт для решения
	vertexes map[point.Point]*vertexData // информация о каждой вершине (расстояние, родитель, статус)
	heap     minHeap.MinHeap             // приоритетная очередь с оценочной функцией f(n)
}

// Solve находит кратчайший путь в лабиринте с использованием алгоритма A*.
//
// Метод реализует алгоритм A* с манхэттенским расстоянием в качестве эвристики.
// Использует функцию оценки f(n) = g(n) + h(n), где g(n) - фактическое расстояние
// от старта, а h(n) - манхэттенское расстояние до цели. Это позволяет направлять
// поиск к целевой точке, делая алгоритм более эффективным чем Дейкстра.
//
// Алгоритм работы:
//  1. Инициализирует все вершины с бесконечным расстоянием
//  2. Добавляет стартовую точку в приоритетную очередь с оценкой 0
//  3. Пока очередь не пуста:
//     - Извлекает вершину с минимальной оценкой f(n)
//     - Для каждого соседа вычисляет g(n) + стоимость + h(n)
//     - Обновляет путь, если найден лучший
//  4. Восстанавливает оптимальный путь по родительским связям
func (as *aStarSolver) Solve(maze *maze.Maze, start, finish point.Point) (path.Path, error) {
	if maze == nil {
		return path.NewPath(), errors.New("maze is nil")
	}
	as.init(maze)
	as.vertexes[start] = &vertexData{dist: 0, visited: false, parent: noParentPoint}
	as.heap.Push(start, 0)
	for !as.heap.IsEmpty() {
		currentPoint := as.heap.Pop()
		if as.vertexes[currentPoint].visited {
			continue
		}
		as.vertexes[currentPoint].visited = true
		if currentPoint == finish {
			break
		}
		neighbors := maze.GetFreeNeighbors(currentPoint)
		for _, neighbor := range neighbors {
			if as.vertexes[neighbor].visited {
				continue
			}
			neighborVertex := as.maze.GetCell(neighbor)
			tentativeDist := as.vertexes[currentPoint].dist + neighborVertex.GetCost()
			if tentativeDist < as.vertexes[neighbor].dist {
				as.vertexes[neighbor].dist = tentativeDist
				as.vertexes[neighbor].parent = currentPoint
				h := as.abs(neighbor.X()-finish.X()) + as.abs(neighbor.Y()-finish.Y())
				as.heap.Push(neighbor, tentativeDist+h)
			}
		}
	}
	if !as.vertexes[finish].visited {
		return path.Path{}, fmt.Errorf("no path from (%d, %d) to (%d, %d)",
			start.X(), start.Y(), finish.X(), finish.Y())
	}
	result := path.NewPath()
	for cur := finish; cur != noParentPoint; cur = as.vertexes[cur].parent {
		result.AddPoint(cur)
	}
	return result, nil
}

// abs возвращает абсолютное значение числа.
//
// Вспомогательная функция для вычисления манхэттенского расстояния,
// которое используется в качестве эвристической функции h(n) в алгоритме A*.
// Манхэттенское расстояние вычисляется как сумма абсолютных разностей
// координат по осям X и Y.
//
// Параметры:
//   - num: число, для которого нужно получить абсолютное значение
//
// Возвращает:
//   - int: абсолютное значение числа (всегда неотрицательное)
func (as *aStarSolver) abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// init инициализирует решатель A* и подготавливает структуры данных.
//
// Метод настраивает все необходимые компоненты для выполнения алгоритма A*:
// создаёт приоритетную очередь для хранения вершин с их оценочными функциями f(n),
// инициализирует карту вершин с бесконечными расстояниями и устанавливает
// начальные значения для всех ячеек лабиринта.
//
// Параметры:
//   - maze: лабиринт для инициализации структур данных
//
// Инициализируемые компоненты:
//   - Приоритетная очередь для выбора вершины с минимальной оценкой f(n) = g(n) + h(n)
//   - Карта вершин с информацией о фактическом расстоянии g(n), родителе и статусе
//   - Все расстояния устанавливаются в бесконечность (inf)
//   - Все вершины помечаются как непосещённые
func (as *aStarSolver) init(maze *maze.Maze) {
	as.maze = maze
	as.heap = minHeap.NewMinHeap()
	as.vertexes = make(map[point.Point]*vertexData, maze.Cols()*maze.Rows())
	for i := range maze.Rows() {
		for j := range maze.Cols() {
			p := point.NewPoint(j, i)
			as.vertexes[p] = &vertexData{dist: inf, visited: false, parent: noParentPoint}
		}
	}
}
