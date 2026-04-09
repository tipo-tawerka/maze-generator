# Структура проекта

```
cmd/app/
  main.go           — точка входа
  cmd/              — cobra-команды (generate, solve)
internal/
  domain/           — доменные модели и интерфейсы
  application/      — use cases (генерация, решение)
  infrastructure/   — IO-адаптеры (рендеринг, чтение/запись файлов)
test/
  cases/            — black-box тест-кейсы
```

## Полезные ссылки

- Алгоритмы генерации лабиринтов: https://habr.com/ru/articles/445378/
- DFS: https://ru.algorithmica.org/cs/graph-traversals
- Dijkstra / BFS: https://ru.algorithmica.org/cs/shortest-paths
- Unicode box-drawing символы: https://www.vidarholen.net/cgi-bin/labyrinth
