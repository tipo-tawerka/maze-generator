// Package main представляет точку входа в приложение для генерации и решения лабиринтов.
//
// Приложение предоставляет CLI интерфейс для работы с лабиринтами, включая:
//   - Генерацию лабиринтов различными алгоритмами (DFS, Prim, Random Walk)
//   - Решение лабиринтов с помощью алгоритмов поиска пути (SPFA, Dijkstra, A*)
//   - Отображение лабиринтов с использованием ASCII или Unicode символов
//   - Сохранение результатов в файлы или вывод в консоль
package main

import (
	"fmt"
	"os"

	"github.com/tipo-tawerka/maze-generator/cmd/app/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
