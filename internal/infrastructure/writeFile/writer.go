// Package writeFile предоставляет реализацию записи лабиринтов в файлы.
//
// Пакет содержит реализацию интерфейса MazeWriter, которая записывает
// строковое представление лабиринтов в текстовые файлы на диске.
package writeFile

import "os"

// FileWriter реализует запись лабиринтов в текстовые файлы.
type FileWriter struct{}

// WriteMaze записывает строковое представление лабиринта в файл.
//
// Параметры:
//   - maze: строковое представление лабиринта для записи
//   - filename: имя файла для записи
//
// Возвращает ошибку при проблемах с созданием или записью в файл.
func (c *FileWriter) WriteMaze(maze, filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	_, err = file.WriteString(maze)
	if err != nil {
		return err
	}
	return nil
}
