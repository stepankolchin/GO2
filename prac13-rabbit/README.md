## Практическое занятие №13 Подключение к RabbitMQ. Отправка и получение сообщений. Колчин Степан Сергеевич ЭФМО-02-25.

## Описание проекта

Реализован асинхронный обмен сообщениями через RabbitMQ между сервисом `tasks` и отдельным worker-процессом.

### Архитектура
```text
Клиент → POST /v1/tasks → tasks-service → RabbitMQ (task_events) → worker → логи
```

### Сценарий работы

1. Клиент отправляет POST-запрос на создание задачи
2. Сервис `tasks` создаёт задачу в in-memory хранилище
3. Сервис публикует событие `task.created` в очередь `task_events`
4. Worker получает сообщение из очереди и логирует его
5. Worker подтверждает обработку через `ack`

---

## Требования

- Go 1.23+
- Docker / Docker Compose
- RabbitMQ 3.x (management)


## Структура проекта
```text
prac13-rabbit/
├── deploy/
│   └── rabbit/
│       └── docker-compose.yml
├── services/
│   ├── tasks/
│   │   └── cmd/
│   │       └── tasks/
│   │           └── main.go
│   └── worker/
│       └── cmd/
│           └── worker/
│               └── main.go
├── internal/
│   ├── amqp/
│   │   └── publisher.go
│   ├── events/
│   │   └── event.go
│   ├── httpapi/
│   │   └── handler.go
│   └── task/
│       ├── model.go
│       └── repo.go
├── go.mod
└── README.md
```
## Команды запуска

### Запуск RabbitMQ

```bash
cd deploy/rabbit
docker compose up -d
```

### Запуск worker
```bash
cd services/worker
export RABBIT_URL=amqp://guest:guest@localhost:5672/
export QUEUE_NAME=task_events
go run ./cmd/worker
```

### Запуск tasks сервиса
```bash
cd services/tasks
export RABBIT_URL=amqp://guest:guest@localhost:5672/
export QUEUE_NAME=task_events
go run ./cmd/tasks
```


## Отчётные материалы

`Запуск docker-compose`

<img width="880" height="788" alt="image" src="https://github.com/user-attachments/assets/9a23c110-bff7-4968-a243-6f973216fa76" />

`management UI http://localhost:15672/`

<img width="974" height="498" alt="image" src="https://github.com/user-attachments/assets/d4f9c2a4-66b7-4dc0-8680-798e665a664d" />

`Отправка запроса`

<img width="880" height="138" alt="image" src="https://github.com/user-attachments/assets/6d3fd47f-af81-4ff1-a9bc-9882f90e0e8b" />

`Проверка логов worker`

<img width="889" height="145" alt="image" src="https://github.com/user-attachments/assets/b98a3ee8-f888-4b81-9ecf-e872c9bd02b9" />


`Queues`

<img width="974" height="466" alt="image" src="https://github.com/user-attachments/assets/5199da5d-ee1d-4186-afaf-0987172d65e6" />
