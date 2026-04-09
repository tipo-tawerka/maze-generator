// Package generating предоставляет функцию для генерации лабиринтов.
//
// Пакет содержит основную логику приложения для создания лабиринтов.
package generating

import (
	"fmt"
	"strconv"

	"github.com/tipo-tawerka/maze-generator/internal/domain/IO"
	"github.com/tipo-tawerka/maze-generator/internal/domain/generator"
	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
)

type GenerateMaze struct {
	Printer IO.MazePrinter
	Writer  IO.MazeWriter
}

// Generate выполняет полный цикл генерации лабиринта.
//
// Параметры:
//   - algorithm: название алгоритма генерации ("dfs", "prim", "randwalk")
//   - output: путь к файлу для сохранения (пустая строка для вывода в консоль)
//   - width: ширина лабиринта в виде строки (должна быть нечетным числом)
//   - height: высота лабиринта в виде строки (должна быть нечетным числом)
//
// Возвращает ошибку в случае проблем с валидацией параметров,
// созданием генератора, генерацией лабиринта или записью результата.
func (gm *GenerateMaze) Generate(algorithm, output, width, height string) error {
	Generator, err := generator.FabricateGenerator(algorithm)
	if err != nil {
		return err
	}
	widthInt, heightInt, err := gm.parseBorders(width, height)
	if err != nil {
		return err
	}
	Maze := maze.NewMaze(widthInt+2, heightInt+2)
	err = Generator.Generate(&Maze)
	if err != nil {
		return err
	}
	mazeStr, err := gm.Printer.PrintMaze(&Maze)
	if err != nil {
		return err
	}
	return gm.Writer.WriteMaze(mazeStr, output)
}

// parseBorders валидирует и парсит строковые значения ширины и высоты лабиринта.
//
// Проверяет, что переданные строки являются корректными числами.
//
// Параметры:
//   - width: ширина лабиринта в виде строки
//   - height: высота лабиринта в виде строки
//
// Возвращает:
//   - int: ширина лабиринта как целое число
//   - int: высота лабиринта как целое число
//   - error: ошибку, если значения невалидны или не являются нечетными
func (gm *GenerateMaze) parseBorders(width, height string) (int, int, error) {
	w, err := strconv.Atoi(width)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid width: %w", err)
	}
	h, err := strconv.Atoi(height)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid height: %w", err)
	}
	if h%2 != 1 {
		return 0, 0, fmt.Errorf("height must be odd")
	}
	if w%2 != 1 {
		return 0, 0, fmt.Errorf("width must be odd")
	}
	return w, h, nil
}
