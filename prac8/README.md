## Практическое занятие №8 Настройка GitHub Actions / GitLab CI для деплоя приложения. Колчин Степан Сергеевич ЭФМО-02-25.

## Тема: 
Настройка GitHub Actions / GitLab CI для деплоя приложения

## Цель работы: 
освоить основы CI/CD для backend-проекта на Go, научиться настраивать автоматический pipeline для проверки, сборки и упаковки Docker-образа

## Краткое объяснение CI и CD

`CI (Continuous Integration)` - непрерывная интеграция
Это автоматический процесс, который запускается при каждом изменении кода в репозитории. Pipeline проверяет, что код собирается, тесты проходят и ничего не сломалось

`CD (Continuous Delivery/Deployment)` - непрерывная доставка/развёртывание.
Это следующий этап: после успешной проверки код автоматически упаковывается в Docker-образ, публикуется в registry и/или разворачивается на сервере

## Структура pipeline

Pipeline состоит из двух job:

- `test-and-build`	Установка Go, загрузка зависимостей, запуск тестов, компиляция

- `docker-build`	Зависит от успеха первого job, сборка Docker-образа

## Выбранная платформа

`GitHub Actions`

## полный YAML-файл pipeline;  
```yaml
name: CI Pipeline

on:
  push:
    branches: [ main, master ]
  pull_request:
    branches: [ main, master ]

jobs:
  test-and-build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.23'

      - name: Show Go version
        run: go version

      - name: Download dependencies
        run: go mod tidy
        working-directory: ./prac7-docker/services/tasks

      - name: Run tests
        run: go test ./...
        working-directory: ./prac7-docker/services/tasks

      - name: Build application
        run: go build ./...
        working-directory: ./prac7-docker/services/tasks

  docker-build:
    runs-on: ubuntu-latest
    needs: test-and-build
    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Build Docker image
        run: docker build -t techip-tasks:${{ github.sha }} .
        working-directory: ./prac7-docker/services/tasks
```

## пояснение шагов pipeline

## Пояснение шагов pipeline

| Шаг | Действие | Назначение |
|-----|----------|------------|
| `actions/checkout@v4` | Клонирование репозитория | Получить код для работы |
| `actions/setup-go@v5` | Установка Go | Настроить окружение для сборки |
| `go version` | Проверка версии Go | Убедиться, что нужная версия установлена |
| `go mod tidy` | Загрузка зависимостей | Скачать все необходимые модули |
| `go test ./...` | Запуск тестов | Проверить корректность кода |
| `go build ./...` | Сборка приложения | Убедиться, что код компилируется |
| `docker/setup-buildx-action@v3` | Настройка Docker Buildx | Подготовить Docker для сборки |
| `docker build` | Сборка образа | Создать Docker-образ с приложением |

## Способ формирования тега образа

```yaml
docker build -t techip-tasks:${{ github.sha }} .
```

В качестве тега используется хеш коммита (github.sha).
Это позволяет:

- Однозначно связать версию образа с конкретным коммитом

- Отследить, какой код соответствует какому образу

- При необходимости откатиться к предыдущей версии

## где должны храниться секреты

Секреты (токены, пароли, SSH-ключи) должны храниться в GitHub Secrets, а не в репозитории или YAML-файле.

Как добавить секрет:

- Репозиторий → Settings → Secrets and variables → Actions

- New repository secret

## скриншот успешного выполнения pipeline

![photo](./photos/image.png)
