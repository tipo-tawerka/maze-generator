package generator

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"time"

	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

// primGenerator реализует алгоритм генерации лабиринта на основе алгоритма Прима.
//
// Алгоритм создает минимальное остовное дерево, что приводит к генерации
// лабиринтов с более равномерным распределением проходов и большим количеством
// коротких тупиков по сравнению с DFS. Процесс начинается с одной ячейки
// и постепенно расширяется, добавляя соседние стены в список кандидатов.
//
// Особенности алгоритма:
//   - Создает более равномерные лабиринты с хорошей связностью
//   - Генерирует множество коротких тупиков
//   - Имеет более предсказуемую структуру по сравнению с DFS
type primGenerator struct {
	maze   *maze.Maze    // указатель на генерируемый лабиринт
	rand   *rand.Rand    // генератор псевдослучайных чисел
	neighs []point.Point // список соседних стен-кандидатов для преобразования в проходы
	dir    [4]direction  // массив направлений движения (вверх, вправо, вниз, влево)
}

// Generate генерирует лабиринт с использованием алгоритма Прима.
//
// Метод реализует модифицированную версию алгоритма Прима для построения
// минимального остовного дерева. Начинает с одной ячейки и постепенно
// расширяется, случайным образом выбирая из списка граничных стен для
// преобразования в проходы. Это создает более равномерную структуру
// лабиринта по сравнению с алгоритмом поиска в глубину.
//
// Алгоритм работы:
//  1. Инициализирует лабиринт стенами и устанавливает стартовую точку
//  2. Добавляет соседние стены стартовой точки в список кандидатов
//  3. Пока список кандидатов не пуст:
//     - Случайно выбирает стену из списка
//     - Преобразует ее в проход
//     - Соединяет с существующими проходами
//     - Добавляет новые соседние стены в список кандидатов
func (pg *primGenerator) Generate(maze *maze.Maze) error {
	if maze == nil {
		return errors.New("maze is nil")
	}
	err := pg.setStartState(maze)
	if err != nil {
		return err
	}
	for len(pg.neighs) > 0 {
		idx := pg.rand.Intn(len(pg.neighs))
		currentPoint := pg.neighs[idx]
		pg.neighs = append(pg.neighs[:idx], pg.neighs[idx+1:]...)

		currentCell, err := cellType.NewCellType(currentPoint, cellType.Empty)
		if err != nil {
			return err
		}
		pg.maze.SetCell(currentCell)

		adjacentEmptyCells := make([]point.Point, 0, 4)
		for _, elem := range pg.dir {
			nPoint := point.NewPoint(currentPoint.X()+elem.x, currentPoint.Y()+elem.y)
			if pg.maze.IsValid(nPoint) {
				neighborCell := pg.maze.GetCell(nPoint)
				if neighborCell.IsEmpty() {
					adjacentEmptyCells = append(adjacentEmptyCells, nPoint)
				}
			}
		}
		if len(adjacentEmptyCells) == 0 {
			return fmt.Errorf("no adjacent empty cells found in maze")
		}
		chosenPoint := adjacentEmptyCells[pg.rand.Intn(len(adjacentEmptyCells))]

		wallX := (currentPoint.X() + chosenPoint.X()) / 2
		wallY := (currentPoint.Y() + chosenPoint.Y()) / 2
		wallCell, err := cellType.NewCellType(point.NewPoint(wallX, wallY), cellType.Empty)
		if err != nil {
			return err
		}
		pg.maze.SetCell(wallCell)
		pg.appendNeighbors(currentPoint)
	}
	return nil
}

// appendNeighbors добавляет соседние стены указанной точки в список кандидатов.
//
// Метод проверяет все четыре направления от текущей точки и добавляет
// соседние ячейки, которые являются стенами, в список потенциальных
// кандидатов для преобразования в проходы. Избегает дублирования, проверяя,
// не находится ли уже соседняя ячейка в списке кандидатов.
func (pg *primGenerator) appendNeighbors(currentPoint point.Point) {
	for _, elem := range pg.dir {
		nPoint := point.NewPoint(currentPoint.X()+elem.x, currentPoint.Y()+elem.y)
		if pg.maze.IsValid(nPoint) {
			neighborCell := pg.maze.GetCell(nPoint)
			if neighborCell.IsWall() {
				found := slices.Contains(pg.neighs, nPoint)
				if !found {
					pg.neighs = append(pg.neighs, nPoint)
				}
			}
		}
	}
}

// setStartState устанавливает начальное состояние для генерации лабиринта алгоритмом Прима.
//
// Метод инициализирует генератор, устанавливает стартовую ячейку (1,1) как пустую
// и добавляет все соседние стены стартовой ячейки в список кандидатов.
func (pg *primGenerator) setStartState(maze *maze.Maze) error {
	err := pg.init(maze)
	if err != nil {
		return err
	}
	startPoint := point.NewPoint(1, 1)
	startCell, err := cellType.NewCellType(startPoint, cellType.Empty)
	if err != nil {
		return err
	}
	pg.maze.SetCell(startCell)
	pg.appendNeighbors(startPoint)
	return nil
}

// init инициализирует генератор Прима и подготавливает лабиринт к генерации.
func (pg *primGenerator) init(maze *maze.Maze) error {
	pg.maze = maze
	pg.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	for y := range maze.Rows() {
		for x := range maze.Cols() {
			cell, err := cellType.NewCellType(point.NewPoint(x, y), cellType.Wall)
			if err != nil {
				return err
			}
			pg.maze.SetCell(cell)
		}
	}
	pg.neighs = make([]point.Point, 0, 32)
	pg.dir = [4]direction{
		{x: 0, y: -2},
		{x: 2, y: 0},
		{x: 0, y: 2},
		{x: -2, y: 0},
	}
	return nil
}
