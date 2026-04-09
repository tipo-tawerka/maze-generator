package generator

import (
	"errors"
	"math/rand"
	"time"

	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

// dfsGenerator реализует алгоритм генерации лабиринта с использованием поиска в глубину (DFS).
//
// Алгоритм создает лабиринты с характерными длинными извилистыми коридорами.
// Процесс генерации начинается с полностью
// заполненного стенами лабиринта и постепенно "прорывает" проходы, используя
// рекурсивный обход в глубину.
//
// Особенности алгоритма:
//   - Создает лабиринты с высокой связностью
//   - Генерирует длинные извилистые пути
type dfsGenerator struct {
	maze *maze.Maze // указатель на генерируемый лабиринт
	rand *rand.Rand // генератор псевдослучайных чисел для случайного выбора направлений
}

// Generate генерирует лабиринт с использованием алгоритма поиска в глубину.
//
// Метод инициализирует лабиринт стенами, устанавливает начальную точку (1,1)
// как пустую ячейку и запускает рекурсивный процесс "прорезания" проходов.
//
// Алгоритм работы:
//  1. Заполняет весь лабиринт стенами
//  2. Устанавливает стартовую точку (1,1) как пустую ячейку
//  3. Рекурсивно прорезает пути от стартовой точки
func (dfs *dfsGenerator) Generate(maze *maze.Maze) error {
	if maze == nil {
		return errors.New("maze is nil")
	}
	err := dfs.init(maze)
	if err != nil {
		return err
	}
	startPoint := point.NewPoint(1, 1)
	startCell, err := cellType.NewCellType(startPoint, cellType.Empty)
	if err != nil {
		return err
	}
	dfs.maze.SetCell(startCell)
	return dfs.carve(startPoint)
}

// init инициализирует генератор DFS и подготавливает лабиринт к генерации.
//
// Метод устанавливает ссылку на лабиринт, инициализирует генератор случайных чисел
// с текущим временем как seed и заполняет весь лабиринт стенами.
func (dfs *dfsGenerator) init(maze *maze.Maze) error {
	dfs.maze = maze
	dfs.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	for y := range maze.Rows() {
		for x := range maze.Cols() {
			cell, err := cellType.NewCellType(point.NewPoint(x, y), cellType.Wall)
			if err != nil {
				return err
			}
			dfs.maze.SetCell(cell)
		}
	}
	return nil
}

// carve рекурсивно прорезает проходы в лабиринте, начиная с указанной точки.
//
// Метод является основой алгоритма DFS. Он случайным образом перемешивает
// направления движения и для каждого направления проверяет возможность
// прорезания прохода. Если соседняя ячейка через одну является стеной,
// создается проход как к промежуточной, так и к целевой ячейке, после
// чего алгоритм рекурсивно вызывается для новой позиции.
func (dfs *dfsGenerator) carve(newPoint point.Point) error {
	directions := dfs.shuffleDirections()
	for _, dir := range directions {
		nx := newPoint.X() + dir.x
		ny := newPoint.Y() + dir.y
		nPoint := point.NewPoint(nx, ny)
		if dfs.maze.IsValid(nPoint) {
			neighborCell := dfs.maze.GetCell(nPoint)
			if neighborCell.IsWall() {
				midX := newPoint.X() + dir.x/2
				midY := newPoint.Y() + dir.y/2
				midCell, err := cellType.NewCellType(point.NewPoint(midX, midY), cellType.Empty)
				if err != nil {
					return err
				}
				dfs.maze.SetCell(midCell)
				neighborEmptyCell, err := cellType.NewCellType(nPoint, cellType.Empty)
				if err != nil {
					return err
				}
				dfs.maze.SetCell(neighborEmptyCell)
				err = dfs.carve(nPoint)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// shuffleDirections перемешивает направления движения для обеспечения случайности генерации.
//
// Метод создает массив из четырех основных направлений (вверх, вправо, вниз, влево)
// с шагом 2 (чтобы оставлять место для стен между проходами) и случайным образом
// перемешивает их. Это гарантирует, что каждый лабиринт будет иметь уникальную
// структуру проходов.
//
// Возвращает:
//   - [4]direction: массив перемешанных направлений движения
//
// Направления:
//   - {0, -2}: движение вверх на 2 ячейки
//   - {2, 0}:  движение вправо на 2 ячейки
//   - {0, 2}:  движение вниз на 2 ячейки
//   - {-2, 0}: движение влево на 2 ячейки
func (dfs *dfsGenerator) shuffleDirections() [4]direction {
	directions := [4]direction{
		{0, -2},
		{2, 0},
		{0, 2},
		{-2, 0},
	}
	dfs.rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})
	return directions
}
