// Package cmd содержит реализацию CLI команд для приложения лабиринтов.
//
// Пакет предоставляет две основные команды:
//   - generate - для генерации лабиринтов различными алгоритмами
//   - solve - для решения лабиринтов с помощью алгоритмов поиска пути
//
// Команды используют библиотеку cobra для обработки аргументов командной строки
// и предоставляют детальную справку по использованию.
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tipo-tawerka/maze-generator/internal/application/generating"
	"github.com/tipo-tawerka/maze-generator/internal/domain/IO"
	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/printChar"
	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/printUnicode"
	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/writeConsole"
	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/writeFile"
)

// generateCmd определяет команду для генерации лабиринтов.
// Поддерживает различные алгоритмы генерации и настройки вывода.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Сгенерировать лабиринт",
	Long: `Генерирует лабиринт с использованием различных алгоритмов.

Доступные алгоритмы:
  • dfs      - поиск в глубину (Depth-First Search)
  • prim     - алгоритм Прима (Prim's algorithm)
  • randwalk - случайное блуждание (Random Walk)

Размеры лабиринта (ширина и высота) должны быть нечетными числами.
Если не указан выходной файл, лабиринт выводится в консоль.`,
	Example: `  # Создать лабиринт 21x21 алгоритмом DFS и сохранить в файл
  app generate --algorithm=dfs --width=21 --height=21 --output=maze.txt

  # Создать маленький лабиринт 11x11 алгоритмом Прима
  app generate --algorithm=prim --width=11 --height=11

  # Создать лабиринт с Unicode символами
  app generate --algorithm=randwalk --width=15 --height=15 --unicode`,
	RunE:          generateRunE,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// generateRunE является функцией выполнения команды generate.
// Настраивает провайдеры вывода и форматирования в зависимости от флагов,
// создает экземпляр генератора и запускает процесс генерации лабиринта.
//
// Параметры:
//   - cmd: команда cobra с установленными флагами
//   - _: неиспользуемые аргументы командной строки
//
// Возвращает ошибку в случае проблем с генерацией или записью лабиринта.
func generateRunE(cmd *cobra.Command, _ []string) error {
	output := cmd.Flag("output").Value.String()
	var outputProvider IO.MazeWriter
	var printerProvider IO.MazePrinter
	if output == "" {
		outputProvider = &writeConsole.ConsoleWriter{}
	} else {
		outputProvider = &writeFile.FileWriter{}
	}
	unicode := cmd.Flags().Changed("unicode")
	if unicode {
		printerProvider = &printUnicode.UnicodePrinter{}
	} else {
		printerProvider = &printChar.CharPrinter{}
	}
	generator := generating.GenerateMaze{
		Printer: printerProvider,
		Writer:  outputProvider,
	}
	err := generator.Generate(
		cmd.Flag("algorithm").Value.String(),
		output,
		cmd.Flag("width").Value.String(),
		cmd.Flag("height").Value.String(),
	)
	return err
}

// init инициализирует команду generate и ее флаги.
// Регистрирует команду в корневой команде приложения и устанавливает
// обязательные и опциональные параметры для генерации лабиринта.
func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("algorithm", "a", "",
		"Алгоритм генерации лабиринта (dfs, prim, randwalk)")
	generateCmd.Flags().StringP("output", "o", "",
		"Путь к выходному файлу (если не указан, выводится в консоль)")
	generateCmd.Flags().String("width", "", "Ширина лабиринта (нечетное число)")
	generateCmd.Flags().String("height", "", "Высота лабиринта (нечетное число)")
	generateCmd.Flags().BoolP("unicode", "u", false,
		"Использовать Unicode символы для отображения лабиринта")
	_ = generateCmd.MarkFlagRequired("algorithm")
	_ = generateCmd.MarkFlagRequired("width")
	_ = generateCmd.MarkFlagRequired("height")
}
