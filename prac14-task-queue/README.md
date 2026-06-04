# Практическое занятие №14 Реализация очереди задач (producer–consumer). Колчин Степан Сергеевич ЭФМО-02-25.

## Описание проекта

Реализована очередь задач по модели producer–consumer с использованием RabbitMQ. Система поддерживает повторные попытки обработки (retries), Dead Letter Queue (DLQ) для проблемных сообщений и идемпотентную обработку.

### Архитектура
```text
Клиент → POST /v1/jobs/process-task → tasks-service → RabbitMQ (task_jobs) → worker → (успех: ack / ошибка: retry / лимит: DLQ)
```
### Структура проекта
```text
prac14-task-queue/
├── deploy/
│   └── rabbit/
│       └── docker-compose.yml
├── internal/
│   ├── amqp/
│   │   ├── publisher.go
│   │   └── setup.go
│   ├── httpapi/
│   │   └── handler.go
│   ├── jobs/
│   │   └── task_job.go
│   └── store/
│       └── processed.go
├── services/
│   ├── tasks/
│   │   └── cmd/tasks/main.go
│   └── worker/
│       └── cmd/worker/main.go
├── go.mod
└── README.md
```

### Очереди

| Очередь | Назначение |
|---------|------------|
| `task_jobs` | Основная очередь задач |
| `task_jobs_dlq` | Dead Letter Queue для проблемных сообщений |

### Механизмы надёжности
1. Retries (повторные попытки)
- Максимальное число попыток: 3

- При ошибке обработки attempt увеличивается

- Сообщение публикуется заново в основную очередь

2. Dead Letter Queue (DLQ)
- При превышении лимита попыток сообщение отправляется в task_jobs_dlq

- Настройка через аргументы очереди:

```go
args := amqp.Table{
    "x-dead-letter-exchange":    "",
    "x-dead-letter-routing-key": "task_jobs_dlq",
}
```
3. Идемпотентность
- Каждое сообщение имеет уникальный message_id

- Worker хранит в памяти map[string]bool обработанных ID

- При повторной доставке сообщение подтверждается без выполнения работы

4. Ручное подтверждение (ack)
- Worker обрабатывает сообщение и явно отправляет ack

- При ошибке — nack или повторная публикация

### Требования
Go 1.23+
Docker / Docker Compose
RabbitMQ 3.x

## Запуск RabbitMQ
```bash
cd deploy/rabbit
docker compose up -d
```
## Запуск worker
```bash
cd services/worker
export RABBIT_URL=amqp://guest:guest@localhost:5672/
go run ./cmd/worker
```

### Запуск tasks сервиса
```bash
cd services/tasks
export RABBIT_URL=amqp://guest:guest@localhost:5672/
go run ./cmd/tasks
```

## Отчетные материалы

`Запуск Docker Compose для RabbitMQ`

<img width="880" height="788" alt="image" src="https://github.com/user-attachments/assets/60b565a8-bdd4-408d-a64d-4efda3230222" />

`Успешная задача`

<img width="874" height="109" alt="image" src="https://github.com/user-attachments/assets/18748ba3-e7c4-4ee2-b9f0-93469a7d5096" />

`Задача с ошибкой`

<img width="866" height="117" alt="image" src="https://github.com/user-attachments/assets/d5545d5b-623d-464f-bf68-402ed86ecd63" />

`Лог worker 3 ошибки - DLQ`

<img width="974" height="224" alt="image" src="https://github.com/user-attachments/assets/60cc689d-bb89-406e-b31c-ae373844dcda" />

`Queues`

<img width="974" height="291" alt="image" src="https://github.com/user-attachments/assets/26f8d0a9-dce7-4b5e-b1b6-e93ee6687149" />
