// Package solving предоставляет функцию для решения лабиринтов.
//
// Пакет содержит основную логику приложения для поиска пути в лабиринтах.
package solving

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tipo-tawerka/maze-generator/internal/domain/IO"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
	"github.com/tipo-tawerka/maze-generator/internal/domain/solver"
)

type SolveMaze struct {
	Reader  IO.MazeReader
	Printer IO.MazePrinter
	Writer  IO.MazeWriter
}

// Solve выполняет полный цикл решения лабиринта.
//
// Параметры:
//   - algorithm: название алгоритма решения ("spfa", "dijkstra", "astar")
//   - input: путь к файлу с лабиринтом для решения
//   - output: путь к файлу для сохранения результата (пустая строка для вывода в консоль)
//   - start: координаты начальной точки в формате "X,Y"
//   - finish: координаты конечной точки в формате "X,Y"
//
// Возвращает ошибку в случае проблем с чтением файла, валидацией координат,
// созданием решателя, поиском пути или записью результата.
func (s *SolveMaze) Solve(algorithm, input, output, start, finish string) error {
	solverAlg, err := solver.FabricateSolver(algorithm)
	if err != nil {
		return err
	}
	Maze, err := s.Reader.ReadMaze(input)
	if err != nil {
		return err
	}
	startPoint, err := s.parseCoordinate(start)
	if err != nil {
		return err
	}
	finishPoint, err := s.parseCoordinate(finish)
	if err != nil {
		return err
	}
	path, err := solverAlg.Solve(&Maze, startPoint, finishPoint)
	if err != nil {
		return err
	}
	err = path.PrintPathOnMaze(Maze)
	if err != nil {
		return err
	}
	mazeStr, err := s.Printer.PrintMaze(&Maze)
	if err != nil {
		return err
	}
	err = s.Writer.WriteMaze(mazeStr, output)
	if err != nil {
		return err
	}
	return nil
}

// parseCoordinate парсит строковое представление координат в объект Point.
//
// Параметры:
//   - coord: строка с координатами в формате "X,Y"
//
// Возвращает:
//   - point.Point: объект с координатами
//   - error: ошибку, если формат неверный или координаты не являются числами
func (s *SolveMaze) parseCoordinate(coord string) (point.Point, error) {
	parts := strings.Split(coord, ",")
	if len(parts) != 2 {
		return point.NewPoint(-1, -1), fmt.Errorf("coordinate should be in format X,Y")
	}
	x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return point.NewPoint(-1, -1), fmt.Errorf("invalid X coordinate: %w", err)
	}
	y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return point.NewPoint(-1, -1), fmt.Errorf("invalid Y coordinate: %w", err)
	}
	return point.NewPoint(x, y), nil
}
