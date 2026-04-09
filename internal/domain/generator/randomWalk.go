package generator

import (
	"errors"
	"math/rand"
	"time"

	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

// PercentOfWalls определяет долю ячеек, которые останутся стенами в лабиринте.
// Значение 0.5 означает, что 50% ячеек будут стенами, а 50% - проходами.
const PercentOfWalls = 0.5

// randomWalkGenerator реализует алгоритм генерации лабиринта методом случайного блуждания.
//
// Алгоритм создает лабиринты с высокой степенью случайности и множественными путями.
// В отличие от DFS и Прима, случайное блуждание может создавать циклы и более
// сложные структуры проходов. Алгоритм также поддерживает создание различных
// типов ячеек (обычные проходы, скоростные дороги, ямы) с разными весами.
//
// Особенности алгоритма:
//   - Создает лабиринты с множественными путями и циклами
//   - Поддерживает различные типы ячеек с разными весами передвижения
//   - Имеет ограничение по времени выполнения (5 секунд)
//   - Генерирует более открытые и менее предсказуемые лабиринты
//   - Подходит для создания игровых уровней с разнообразием путей
type randomWalkGenerator struct {
	maze         *maze.Maze   // указатель на генерируемый лабиринт
	rand         *rand.Rand   // генератор псевдослучайных чисел
	dir          [4]direction // массив направлений движения
	targetCells  int          // целевое количество ячеек-проходов
	createdCells int          // текущее количество созданных ячеек-проходов
}

// Generate генерирует лабиринт с использованием алгоритма случайного блуждания.
//
// Метод запускает генерацию в отдельной горутине с таймаутом 5 секунд.
// Это необходимо, поскольку случайное блуждание может потенциально работать
// бесконечно долго в некоторых конфигурациях. Алгоритм создает проходы
// случайным образом, что может привести к созданию циклов и множественных путей.
func (rw *randomWalkGenerator) Generate(maze *maze.Maze) error {
	if maze == nil {
		return errors.New("maze is nil")
	}
	errChan := make(chan error)
	go rw.generate(maze, errChan)
	select {
	case err := <-errChan:
		return err
	case <-time.After(5 * time.Second):
		return nil
	}
}

// generate выполняет основную логику генерации лабиринта случайным блужданием.
//
// Метод реализует алгоритм случайного блуждания, который начинает с одной точки
// и случайным образом расширяется во всех направлениях. В отличие от DFS и Прима,
// этот алгоритм может посещать уже созданные проходы, что приводит к образованию
// циклов. Генерация продолжается до достижения целевого количества ячеек-проходов.
func (rw *randomWalkGenerator) generate(maze *maze.Maze, errChan chan<- error) {
	err := rw.init(maze)
	if err != nil {
		errChan <- err
		return
	}
	rw.targetCells = int(float64(maze.Rows()*maze.Cols()) * (1 - PercentOfWalls))
	curPoint := point.NewPoint(1, 1)
	startCell, err := rw.newEmptyCell(curPoint)
	if err != nil {
		errChan <- err
		return
	}
	rw.maze.SetCell(startCell)
	visitedDirs := make([]point.Point, 0, 4)
	for rw.targetCells > rw.createdCells {
		for _, elem := range rw.dir {
			nx := curPoint.X() + elem.x
			ny := curPoint.Y() + elem.y
			mx := nx - elem.x/2
			my := ny - elem.y/2
			nPoint := point.NewPoint(nx, ny)
			mPoint := point.NewPoint(mx, my)
			if rw.maze.IsValid(nPoint) {
				mCell := rw.maze.GetCell(mPoint)
				if mCell.IsWall() {
					mEmptyCell, err := rw.newEmptyCell(mPoint)
					if err != nil {
						errChan <- err
						return
					}
					rw.maze.SetCell(mEmptyCell)
				}
				nCell := rw.maze.GetCell(nPoint)
				if nCell.IsWall() {
					nEmptyCell, err := rw.newEmptyCell(nPoint)
					if err != nil {
						errChan <- err
						return
					}
					rw.maze.SetCell(nEmptyCell)
				}
				visitedDirs = append(visitedDirs, nPoint)
			}
		}
		if len(visitedDirs) == 0 {
			break
		}
		curPoint = visitedDirs[rw.rand.Intn(len(visitedDirs))]
	}
	errChan <- nil
}

// init инициализирует генератор случайного блуждания и подготавливает лабиринт.
func (rw *randomWalkGenerator) init(maze *maze.Maze) error {
	rw.maze = maze
	rw.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	rw.dir = [4]direction{
		{x: 0, y: -2},
		{x: 2, y: 0},
		{x: 0, y: 2},
		{x: -2, y: 0},
	}
	for y := range maze.Rows() {
		for x := range maze.Cols() {
			cell, err := cellType.NewCellType(point.NewPoint(x, y), cellType.Wall)
			if err != nil {
				return err
			}
			rw.maze.SetCell(cell)
		}
	}
	return nil
}

// newEmptyCell создает новую ячейку-проход случайного типа.
// // Тип ячейки выбирается случайным образом с заданными вероятностями:
// //   - 60% вероятность создания обычной пустой ячейки
// //   - 20% вероятность создания ячейки скоростной дороги
// //   - 20% вероятность создания ячейки ямы
// Примечание: При успешном создании ячейки увеличивается счетчик созданных ячеек.
func (rw *randomWalkGenerator) newEmptyCell(point point.Point) (cellType.CellType, error) {
	num := rw.rand.Float32()
	typeCell := cellType.Empty
	switch {
	case num < 0.6:
	case num < 0.8:
		typeCell = cellType.Highway
	default:
		typeCell = cellType.Pits
	}
	cell, err := cellType.NewCellType(point, typeCell)
	if err == nil {
		rw.createdCells++
	}
	return cell, err
}
