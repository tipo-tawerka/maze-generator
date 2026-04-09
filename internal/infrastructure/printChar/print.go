// Package printChar предоставляет реализацию печати лабиринтов с использованием ASCII символов.
//
// Пакет содержит реализацию интерфейса MazePrinter, которая форматирует лабиринты
// в строки, используя стандартные ASCII символы для отображения различных типов ячеек.
package printChar

import (
	"errors"

	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

// CharPrinter реализует печать лабиринтов с использованием ASCII символов.
type CharPrinter struct{}

// PrintMaze преобразует лабиринт в строковое представление с ASCII символами.
//
// Параметры:
//   - maze: лабиринт для преобразования в строку
//
// Возвращает:
//   - string: строковое представление лабиринта с символами ASCII
//   - error: всегда nil для данной реализации
func (cp *CharPrinter) PrintMaze(maze *maze.Maze) (string, error) {
	if maze == nil {
		return "", errors.New("maze is nil")
	}
	result := make([]byte, 0, 512)
	for i := range maze.Rows() {
		for j := range maze.Cols() {
			cell := maze.GetCell(point.NewPoint(j, i))
			result = append(result, cell.Print())
		}
		result = append(result, '\n')
	}
	return string(result), nil
}
