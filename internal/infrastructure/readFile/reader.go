// Package readFile предоставляет реализацию чтения лабиринтов из файлов.
//
// Пакет содержит реализацию интерфейса MazeReader, которая читает лабиринты
// из текстовых файлов.
package readFile

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

// FileReader реализует чтение лабиринтов из текстовых файлов.
type FileReader struct{}

// errInvalidFileFormat определяет ошибку неверного формата файла.
var errInvalidFileFormat = errors.New("invalid file format")

// ReadMaze читает лабиринт из указанного текстового файла.
//
// Параметры:
//   - filename: путь к файлу с лабиринтом
//
// Возвращает:
//   - maze.Maze: структуру лабиринта
//   - error: ошибку при проблемах с чтением или парсингом файла
func (fr *FileReader) ReadMaze(filename string) (maze.Maze, error) {
	file, err := os.Open(filename)
	if err != nil {
		return maze.NewMaze(0, 0), err
	}
	defer func() {
		_ = file.Close()
	}()
	reader := bufio.NewReader(file)
	mazeLines := make([][]byte, 0, 16)
	for {
		data, err := reader.ReadBytes('\n')
		switch {
		case err == nil:
			if len(data) == 0 {
				return maze.NewMaze(0, 0),
					fmt.Errorf("%w: пустая строка в файле", errInvalidFileFormat)
			}
			data = bytes.TrimRight(data, "\r\n")
			newData := make([]byte, len(data))
			copy(newData, data)
			mazeLines = append(mazeLines, newData)
		case !errors.Is(err, io.EOF):
			return maze.NewMaze(0, 0), err
		default:
			if len(data) > 0 && data[0] != '\n' {
				data = bytes.TrimRight(data, "\r\n")
				mazeLines = append(mazeLines, data)
			}
			return fr.parseMaze(mazeLines)
		}
	}
}

// parseMaze преобразует строки файла в структуру лабиринта.
//
// Параметры:
//   - mazeLines: срез строк файла как байтовые массивы
//
// Возвращает структуру лабиринта или ошибку при проблемах с форматом.
func (fr *FileReader) parseMaze(mazeLines [][]byte) (maze.Maze, error) {
	if len(mazeLines) == 0 {
		return maze.NewMaze(0, 0), fmt.Errorf("%w: пустой файл", errInvalidFileFormat)
	}
	countCols := len(mazeLines[0])
	Maze := maze.NewMaze(len(mazeLines), countCols)
	for i, line := range mazeLines {
		if len(line) != countCols {
			return Maze,
				fmt.Errorf("%w: строка %d имеет некорректную длину", errInvalidFileFormat, i+1)
		}
		for j, char := range line {
			cell, err := cellType.NewCellType(point.NewPoint(j, i), char)
			if err != nil {
				return Maze, err
			}
			Maze.SetCell(cell)
		}
	}
	return Maze, nil
}
