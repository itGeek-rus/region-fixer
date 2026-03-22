# Region Fixer
Веб приложение для генерации жалоб на инфраструктуру.

| region-fixer/                                    |
|--------------------------------------------------|
| cmd/ - Точка входа приложения                    |
| internal/ - Внутренние пакеты                    |
| config/ - Конфигурация                           |
| handlers/ - HTTP handlers                        |
| infrastructure/ - Инфраструктура (БД, хранилище) |
| middleware/ - Middleware                         |
| models/ - Модели данных                          |
| repository/ - Интерфейсы репозиториев            |
| router/ - Роутинг                                |
| service/ - Бизнес-логика                         |
| web/ - Frontend (шаблоны, статика)               |
| migrations/ - Миграции БД                        |

# Требования

- Go 1.26+
- Docker и Docker Compose

# Установка

### Клонировать репозиторий

- git clone <url>

### Перейти в директорию

- cd region-fixer

## Установить зависимости

- go mod download

# Taskfile

Проект использует [Task](https://taslfile.dev) - общие команды для локальной разработки и проверки перед Pull Request

### Основные команды

| Команда                   | Назначение    |
|---------------------------|---------------|
 | task / task --list        | Список задач  |
  | task deps                 | go mod tidy |
   | task tools: intasll       | gosec, gocritic, golangci-lint |
    | task dc:up / task dc:down | Docker compose запустить/откатить |
     | task local-run            | Освободить порт и запустить приложение локально |
      | task local-run-bin        | Сборка и запуск бинарника |
       | task build                | Только сборка |
        | task test/ task test:race | Запуск тестов |
         | task cleancode            | gofmt + go vet |
          | task fmt:check            | Проверка формата без правок (для CI) |
           | task lint:all             | cleancode -> gosec -> gocritic -> golangci-lint -> тесты |
            | task health               | curl на /health |
             | task env:print            | Вывод DB_*, SERVER_*, APP_ENV |
              | task free-port            | Убить процесс на порту (macOS/Linux, lsof)

### Установка Task

- macOS: `brew install go-task`
- Другие OS: [документация Task](https://taskfile.dev/installation/)

После `task tools:install` бинарники лежат в `$(go env GOPATH)/bin`. Добавить в каталог в `PATH`, иначе не найдутся `gosec`, `gocritic, `golangci-lint`:

```bash
export PATH="$(go env GOPATH)/bin:$PATH"