// Package path определяет пути в лабиринте и операции с ними.
//
// Пакет предоставляет структуру Path для представления найденного пути
// от начальной до конечной точки в лабиринте. Включает методы для добавления
// точек в путь и отображения пути на лабиринте с соответствующими метками.
package path

import (
	"github.com/tipo-tawerka/maze-generator/internal/domain/cellType"
	"github.com/tipo-tawerka/maze-generator/internal/domain/maze"
	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

// Path представляет путь в лабиринте как последовательность точек.
// Хранит упорядоченный список координат от начальной до конечной точки.
type Path struct {
	points []point.Point
}

// NewPath создает новый пустой путь.
//
// Возвращает новый экземпляр Path.
func NewPath() Path {
	return Path{
		points: make([]point.Point, 0, 16),
	}
}

// AddPoint добавляет точку к пути.
//
// Параметры:
//   - pt: координаты точки для добавления в путь
func (p *Path) AddPoint(pt point.Point) {
	p.points = append(p.points, pt)
}

// PrintPathOnMaze отображает путь на лабиринте, помечая соответствующие ячейки.
//
// Метод отмечает все точки пути символом пути ('.'), а также устанавливает
// специальные символы для начальной ('S') и конечной ('F') точек.
// Предполагается, что первая точка в пути - это финиш, а последняя - старт.
//
// Параметры:
//   - maze: лабиринт, на котором нужно отобразить путь
//
// Возвращает ошибку, если не удается создать ячейки нужного типа.
// Если путь пустой, возвращает nil без изменения лабиринта.
func (p *Path) PrintPathOnMaze(maze maze.Maze) error {
	for _, pt := range p.points {
		cell, err := cellType.NewCellType(pt, cellType.Path)
		if err != nil {
			return err
		}
		maze.SetCell(cell)
	}
	if len(p.points) == 0 {
		return nil
	}
	start, err := cellType.NewCellType(p.points[len(p.points)-1], cellType.Start)
	if err != nil {
		return err
	}
	maze.SetCell(start)
	finish, err := cellType.NewCellType(p.points[0], cellType.Finish)
	if err != nil {
		return err
	}
	maze.SetCell(finish)
	return nil
}
