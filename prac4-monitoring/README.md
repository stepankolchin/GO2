## Практическое занятие №4 Настройка Prometheus + Grafana для метрик. Интеграция с приложением. Колчин Степан Сергеевич ЭФМО-02-25.

## Описание проекта

Реализовано HTTP-приложение на Go с экспортом метрик в формате Prometheus. Настроен сбор метрик с помощью Prometheus и визуализация через Grafana.

### Эндпоинты приложения

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET | `/health` | Проверка работоспособности |
| GET | `/students/{id}` | Получение информации о студенте |
| GET | `/metrics` | Метрики приложения в формате Prometheus |

### Метрики приложения

| Имя метрики | Тип | Описание |
|-------------|-----|----------|
| `app_http_requests_total` | Counter | Общее число HTTP-запросов (лейблы: method, path) |
| `app_http_errors_total` | Counter | Число ошибок (лейблы: method, path, status_code) |
| `app_http_request_duration_seconds` | Histogram | Длительность обработки запроса (лейблы: method, path) |

## Требования

- Go 1.16+
- Prometheus (скачан и установлен)
- Grafana (скачана и установлена)

### Запуск Go-приложения

```bash
cd prac4-monitoring
go run ./cmd/server
```

### Запуск Prometheus

```bash
cd /path/to/grafana/bin
./grafana.exe server
```

## Структура проекта

```text
prac4-monitoring/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── httpapi/
│   │   ├── handler.go
│   │   ├── middleware.go
│   │   └── response_writer.go
│   ├── metrics/
│   │   └── metrics.go
│   └── student/
│       ├── model.go
│       └── repo.go
├── monitoring/
│   └── prometheus.yml
├── photos/                  
├── go.mod
├── go.sum
└── README.md
```
## Отчётные материалы

`Проверка метрик в Prometheus`

<img width="847" height="164" alt="image" src="https://github.com/user-attachments/assets/95e2d0ba-4a64-49d9-9744-bc68e4a60da9" />


`Проверка Targets`

<img width="974" height="186" alt="image" src="https://github.com/user-attachments/assets/cfca5747-f562-4369-ac3b-39f8d9c177b7" />


`успешное добавления источника данных Prometheus`

<img width="974" height="124" alt="image" src="https://github.com/user-attachments/assets/b47f7fe6-f190-4694-9a6d-e64f3c19fda5" />

`Полученный дашборд`

<img width="974" height="461" alt="image" src="https://github.com/user-attachments/assets/0e8e5e9d-8abc-4d88-903a-4529bc603125" />
