// Package printUnicode предоставляет реализацию печати лабиринтов с использованием Unicode символов.
//
// Пакет содержит реализацию интерфейса MazePrinter, которая форматирует лабиринты
// в строки, используя Unicode символы.
package printUnicode

import (
	"errors"

	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

// Константы определяют Unicode символы для различных типов соединений стен.
const (
	UnicodeWallHorizontal    rune = '─'
	UnicodeWallVertical      rune = '│'
	UnicodeWallTopLeft       rune = '┌'
	UnicodeWallTopRight      rune = '┐'
	UnicodeWallBottomLeft    rune = '└'
	UnicodeWallBottomRight   rune = '┘'
	UnicodeWallCrossing      rune = '┼'
	UnicodeWallTeeTop        rune = '┬'
	UnicodeWallTeeBottom     rune = '┴'
	UnicodeWallTeeLeft       rune = '├'
	UnicodeWallTeeRight      rune = '┤'
	UnicodeWallVertSmallDown rune = '╷'
	UnicodeWallVertSmallUp   rune = '╵'
	UnicodeWallTeeSmallLeft  rune = '⊢'
	UnicodeWallTeeSmallRight rune = '⊣'
)

// UnicodePrinter реализует печать лабиринтов с использованием Unicode символов.
// Автоматически выбирает подходящие символы соединений в зависимости от
// расположения соседних стен.
type UnicodePrinter struct {
	maze *maze.Maze
}

// PrintMaze преобразует лабиринт в строковое представление с Unicode символами.
//
// Параметры:
//   - maze: лабиринт для преобразования в строку
//
// Возвращает:
//   - string: строковое представление лабиринта с Unicode символами
//   - error: всегда nil для данной реализации
func (cp *UnicodePrinter) PrintMaze(maze *maze.Maze) (string, error) {
	if maze == nil {
		return "", errors.New("maze is nil")
	}
	cp.maze = maze
	result := make([]rune, 0, 512)
	for y := 0; y < maze.Rows(); y += 2 {
		for x := 0; x < maze.Cols(); x++ {
			result = append(result, cp.getUnicode(x, y))
		}
		result = append(result, '\n')
	}
	return string(result), nil
}

// getUnicode определяет подходящий Unicode символ для ячейки в указанных координатах.
//
// Анализирует текущую ячейку и ее соседей, чтобы выбрать наиболее подходящий
// Unicode символ для отображения стены.
//
// Параметры:
//   - x, y: координаты ячейки для анализа
//
// Возвращает подходящий Unicode символ или пробел для проходимых ячеек.
func (cp *UnicodePrinter) getUnicode(x, y int) rune {
	curCell := cp.maze.GetCell(point.NewPoint(x, y))

	if !curCell.IsWall() {
		return ' '
	}

	pointLeftCenter := point.NewPoint(x-1, y)
	pointCenterUp := point.NewPoint(x, y-1)
	pointCenterBottom := point.NewPoint(x, y+1)
	pointRightCenter := point.NewPoint(x+1, y)

	var isWallLeftCenter, isWallCenterUp, isWallCenterBottom, isWallRightCenter bool

	if cp.maze.IsValid(pointLeftCenter) {
		cellLeftCenter := cp.maze.GetCell(pointLeftCenter)
		isWallLeftCenter = cellLeftCenter.IsWall()
	} else {
		isWallLeftCenter = cp.maze.IsLeftBorder(pointLeftCenter)
	}

	if cp.maze.IsValid(pointCenterUp) {
		cellCenterUp := cp.maze.GetCell(pointCenterUp)
		isWallCenterUp = cellCenterUp.IsWall()
	} else {
		isWallCenterUp = cp.maze.IsTopBorder(pointCenterUp)
	}

	if cp.maze.IsValid(pointCenterBottom) {
		cellCenterBottom := cp.maze.GetCell(pointCenterBottom)
		isWallCenterBottom = cellCenterBottom.IsWall()
	} else {
		isWallCenterBottom = cp.maze.IsBottomBorder(pointCenterBottom)
	}

	if cp.maze.IsValid(pointRightCenter) {
		cellRightCenter := cp.maze.GetCell(pointRightCenter)
		isWallRightCenter = cellRightCenter.IsWall()
	} else {
		isWallRightCenter = cp.maze.IsRightBorder(pointRightCenter)
	}
	switch {
	case isWallLeftCenter && isWallRightCenter && isWallCenterUp && isWallCenterBottom:
		return UnicodeWallCrossing
	case isWallLeftCenter && isWallRightCenter && isWallCenterUp:
		return UnicodeWallTeeBottom
	case isWallLeftCenter && isWallRightCenter && isWallCenterBottom:
		return UnicodeWallTeeTop
	case isWallCenterUp && isWallCenterBottom && isWallRightCenter:
		return UnicodeWallTeeLeft
	case isWallCenterUp && isWallCenterBottom && isWallLeftCenter:
		return UnicodeWallTeeRight
	case isWallLeftCenter && isWallRightCenter:
		return UnicodeWallHorizontal
	case isWallCenterUp && isWallCenterBottom:
		return UnicodeWallVertical
	case isWallRightCenter && isWallCenterBottom:
		return UnicodeWallTopLeft
	case isWallLeftCenter && isWallCenterBottom:
		return UnicodeWallTopRight
	case isWallRightCenter && isWallCenterUp:
		return UnicodeWallBottomLeft
	case isWallLeftCenter && isWallCenterUp:
		return UnicodeWallBottomRight
	case isWallLeftCenter && !isWallCenterUp && !isWallCenterBottom && !isWallRightCenter:
		return UnicodeWallTeeSmallRight
	case isWallRightCenter && !isWallCenterUp && !isWallCenterBottom && !isWallLeftCenter:
		return UnicodeWallTeeSmallLeft
	case isWallCenterUp && !isWallLeftCenter && !isWallCenterBottom && !isWallRightCenter:
		return UnicodeWallVertSmallUp
	case isWallCenterBottom && !isWallLeftCenter && !isWallCenterUp && !isWallRightCenter:
		return UnicodeWallVertSmallDown
	default:
		return '+'
	}
}
