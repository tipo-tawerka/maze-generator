// Package IO определяет интерфейсы для ввода/вывода лабиринтов.
package IO

import "github.com/tipo-tawerka/maze-generator/internal/domain/maze"

// MazeReader определяет интерфейс для чтения лабиринтов из внешних источников.
type MazeReader interface {
	// ReadMaze читает лабиринт из указанного источника.
	// Возвращает структуру лабиринта или ошибку при проблемах с чтением.
	ReadMaze(filename string) (maze.Maze, error)
}

// MazePrinter определяет интерфейс для форматирования лабиринтов в строки.
type MazePrinter interface {
	// PrintMaze преобразует лабиринт в строковое представление для отображения.
	// Возвращает отформатированную строку или ошибку при проблемах с форматированием.
	PrintMaze(maze *maze.Maze) (string, error)
}

// MazeWriter определяет интерфейс для записи лабиринтов в различные места назначения.
type MazeWriter interface {
	// WriteMaze записывает строковое представление лабиринта в указанное место.
	// Параметр target определяет место назначения (путь к файлу или другой идентификатор).
	WriteMaze(maze, target string) error
}
