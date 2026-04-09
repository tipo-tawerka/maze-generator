// Package maze предоставляет структуру и операции для работы с лабиринтами.
//
// Пакет содержит основную структуру Maze, которая представляет двумерную сетку
// ячеек лабиринта и предоставляет методы для навигации, проверки границ,
// получения соседей и других операций, необходимых для генерации и решения лабиринтов.
package maze

import (
	"fmt"

	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

// direction представляет направление движения в лабиринте.
// Используется для определения соседних ячеек.
type direction struct {
	x, y int // смещение по координатам x и y
}

// Maze представляет лабиринт как двумерную сетку ячеек.
// Содержит информацию о размерах лабиринта и состоянии каждой ячейки.
type Maze struct {
	maze       [][]cellType.CellType
	rows, cols int
}

// NewMaze создает новый лабиринт с указанными размерами.
//
// Параметры:
//   - rows: количество строк в лабиринте
//   - cols: количество столбцов в лабиринте
//
// Возвращает новый экземпляр Maze с инициализированной сеткой ячеек.
func NewMaze(rows, cols int) Maze {
	maze := make([][]cellType.CellType, rows)
	for i := range maze {
		maze[i] = make([]cellType.CellType, cols)
	}
	return Maze{maze: maze, rows: rows, cols: cols}
}

// GetCell возвращает ячейку лабиринта по указанным координатам.
//
// Параметры:
//   - point: координаты ячейки для получения
//
// Возвращает ячейку по указанным координатам.
// Паникует, если координаты выходят за границы лабиринта.
func (m *Maze) GetCell(point point.Point) cellType.CellType {
	x := point.X()
	y := point.Y()
	if !m.IsValid(point) {
		panic(fmt.Errorf("неверные координаты точки: %d %d, "+
			"когда размеры лабиринта %d %d", x, y, m.cols, m.rows))
	}
	return m.maze[y][x]
}

// SetCell устанавливает ячейку в лабиринте по координатам, указанным в ячейке.
//
// Параметры:
//   - cell: ячейка с координатами и типом для установки в лабиринт
//
// Паникует, если координаты ячейки выходят за границы лабиринта.
func (m *Maze) SetCell(cell cellType.CellType) {
	x := cell.Point.X()
	y := cell.Point.Y()
	if !m.IsValid(cell.Point) {
		panic(fmt.Errorf("неверные координаты точки: %d %d, "+
			"когда размеры лабиринта %d %d", x, y, m.cols, m.rows))
	}
	m.maze[y][x] = cell
}

// Rows возвращает количество строк в лабиринте.
func (m *Maze) Rows() int {
	return m.rows
}

// Cols возвращает количество столбцов в лабиринте.
func (m *Maze) Cols() int {
	return m.cols
}

// IsValid проверяет, находятся ли координаты точки в пределах лабиринта.
//
// Параметры:
//   - point: координаты для проверки
//
// Возвращает true, если координаты находятся в пределах лабиринта.
func (m *Maze) IsValid(point point.Point) bool {
	x := point.X()
	y := point.Y()
	if x < 0 || x >= m.cols || y < 0 || y >= m.rows {
		return false
	}
	return true
}

// GetFreeNeighbors возвращает список проходимых соседних ячеек для указанной точки.
//
// Параметры:
//   - p: координаты точки, для которой ищутся соседи
//
// Возвращает срез координат соседних проходимых ячеек в порядке:
// вверх, вправо, вниз, влево.
func (m *Maze) GetFreeNeighbors(p point.Point) []point.Point {
	directions := []direction{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
	neighbors := make([]point.Point, 0, 4)
	for _, dir := range directions {
		neighbor := point.NewPoint(p.X()+dir.x, p.Y()+dir.y)
		if m.IsValid(neighbor) {
			cell := m.GetCell(neighbor)
			if cell.IsEmpty() {
				neighbors = append(neighbors, neighbor)
			}
		}
	}
	return neighbors
}

// IsLeftBorder проверяет, находится ли точка на левой границе лабиринта.
func (m *Maze) IsLeftBorder(p point.Point) bool {
	return p.X() == 0
}

// IsRightBorder проверяет, находится ли точка на правой границе лабиринта.
func (m *Maze) IsRightBorder(p point.Point) bool {
	return p.X() == m.cols-1
}

// IsTopBorder проверяет, находится ли точка на верхней границе лабиринта.
func (m *Maze) IsTopBorder(p point.Point) bool {
	return p.Y() == 0
}

// IsBottomBorder проверяет, находится ли точка на нижней границе лабиринта.
func (m *Maze) IsBottomBorder(p point.Point) bool {
	return p.Y() == m.rows-1
}
