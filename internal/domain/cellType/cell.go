// Package cellType определяет типы ячеек лабиринта и операции с ними.
//
// Пакет предоставляет различные типы ячеек лабиринта, такие как стены, пустые пространства,
// стартовые и финишные точки, а также специальные поверхности с различной стоимостью
// прохождения. Каждая ячейка имеет координаты и тип, определяющий
// ее поведение в алгоритмах генерации и решения лабиринтов.
package cellType

import (
	"errors"

	"github.com/tipo-tawerka/maze-generator/internal/domain/point"
)

// Константы определяют символьные представления различных типов ячеек лабиринта.
const (
	Wall    byte = '#'
	Empty   byte = ' '
	Start   byte = 'S'
	Finish  byte = 'F'
	Path    byte = '.'
	Pits    byte = 'O'
	Highway byte = '!'
)

// CellType представляет ячейку лабиринта с координатами и типом.
type CellType struct {
	Point    point.Point
	typeCell byte
}

// NewCellType создает новую ячейку с указанными координатами и типом.
//
// Параметры:
//   - p: координаты ячейки в лабиринте
//   - char: символ, определяющий тип ячейки
//
// Возвращает:
//   - CellType: созданную ячейку
//   - error: ошибку, если передан недопустимый символ типа ячейки.
func NewCellType(p point.Point, char byte) (CellType, error) {
	switch char {
	case Wall:
		return CellType{Point: p, typeCell: Wall}, nil
	case Empty:
		return CellType{Point: p, typeCell: Empty}, nil
	case Start:
		return CellType{Point: p, typeCell: Start}, nil
	case Finish:
		return CellType{Point: p, typeCell: Finish}, nil
	case Path:
		return CellType{Point: p, typeCell: Path}, nil
	case Pits:
		return CellType{Point: p, typeCell: Pits}, nil
	case Highway:
		return CellType{Point: p, typeCell: Highway}, nil
	default:
		return CellType{}, errors.New("invalid char")
	}
}

// Print возвращает символьное представление ячейки для отображения.
// Используется при выводе лабиринта в консоль или файл.
func (c *CellType) Print() byte {
	return c.typeCell
}

// IsWall проверяет, является ли ячейка стеной.
// Возвращает true, если ячейка непроходима (стена).
func (c *CellType) IsWall() bool {
	return c.typeCell == Wall
}

// IsEmpty проверяет, является ли ячейка проходимой.
// Возвращает true для всех ячеек, кроме стен.
func (c *CellType) IsEmpty() bool {
	return c.typeCell != Wall
}

// GetCost возвращает стоимость прохождения через ячейку.
// Используется алгоритмами поиска пути для вычисления оптимального маршрута.
//
// Стоимости:
//   - Highway (автострада): 1 - быстрое перемещение
//   - Pits (ямы): 3 - замедленное перемещение
//   - Остальные проходимые ячейки: 2 - обычная скорость
//
// Паникует, если вызывается для стены, так как стены непроходимы.
func (c *CellType) GetCost() int {
	switch c.typeCell {
	case Wall:
		panic("Wall has no cost")
	case Highway:
		return 1
	case Pits:
		return 3
	default:
		return 2
	}
}
