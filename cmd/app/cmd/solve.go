package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tipo-tawerka/maze-generator/internal/application/solving"
	"github.com/tipo-tawerka/maze-generator/internal/domain/IO"
	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/printChar"
	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/readFile"
	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/writeConsole"
	"github.com/tipo-tawerka/maze-generator/internal/infrastructure/writeFile"
)

// solveCmd определяет команду для решения лабиринтов.
// Поддерживает различные алгоритмы поиска пути и настройки ввода/вывода.
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Решить лабиринт",
	Long: `Решить лабиринт, используя различные алгоритмы поиска пути.

Доступные алгоритмы:
  • spfa     - алгоритм SPFA (Shortest Path Faster Algorithm)
  • dijkstra - алгоритм Дейкстры (Dijkstra's algorithm)
  • astar    - алгоритм A* (A-star algorithm)

Команда принимает файл с лабиринтом и находит путь от начальной до конечной точки
с помощью выбранного алгоритма. Результат может быть выведен в консоль или сохранен в файл.`,
	Example: `  # Решить лабиринт алгоритмом SPFA
  app solve --algorithm=spfa --input=maze.txt --start=0,0 --finish=9,9

  # Решить лабиринт и сохранить результат в файл
  app solve --algorithm=astar --input=maze.txt --start=1,1 --finish=19,19 --output=solution.txt

  # Решить лабиринт алгоритмом Дейкстры
  app solve --algorithm=dijkstra --input=maze.txt --start=0,0 --finish=9,9`,
	RunE:          solveRunE,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// solveRunE является функцией выполнения команды solve.
// Настраивает провайдеры ввода/вывода, создает экземпляр решателя
// и запускает процесс поиска пути в лабиринте.
//
// Параметры:
//   - cmd: команда cobra с установленными флагами
//   - _: неиспользуемые аргументы командной строки
//
// Возвращает ошибку в случае проблем с чтением файла, поиском пути или записью результата.
func solveRunE(cmd *cobra.Command, _ []string) error {
	output := cmd.Flag("output").Value.String()
	var outputProvider IO.MazeWriter
	if output == "" {
		outputProvider = &writeConsole.ConsoleWriter{}
	} else {
		outputProvider = &writeFile.FileWriter{}
	}
	solver := solving.SolveMaze{
		Writer:  outputProvider,
		Reader:  &readFile.FileReader{},
		Printer: &printChar.CharPrinter{},
	}
	err := solver.Solve(
		cmd.Flag("algorithm").Value.String(),
		cmd.Flag("file").Value.String(),
		output,
		cmd.Flag("start").Value.String(),
		cmd.Flag("end").Value.String(),
	)
	if err != nil {
		return err
	}
	return nil
}

// init инициализирует команду solve и ее флаги.
// Регистрирует команду в корневой команде приложения и устанавливает
// обязательные и опциональные параметры для решения лабиринта.
func init() {
	rootCmd.AddCommand(solveCmd)
	solveCmd.Flags().StringP("algorithm", "a", "",
		"Алгоритм решения лабиринта (spfa, dijkstra, astar)")
	solveCmd.Flags().StringP("file", "f", "", "Путь к входному файлу с лабиринтом")
	solveCmd.Flags().StringP("output", "o", "",
		"Путь к выходному файлу (если не указан, выводится в консоль)")
	solveCmd.Flags().StringP("start", "s", "", "Начальная координата в формате X,Y")
	solveCmd.Flags().StringP("end", "e", "", "Конечная координата в формате X,Y")
	_ = solveCmd.MarkFlagRequired("algorithm")
	_ = solveCmd.MarkFlagRequired("file")
	_ = solveCmd.MarkFlagRequired("start")
	_ = solveCmd.MarkFlagRequired("end")

}
