## Практическое занятие №7 Написание Dockerfile и сборка контейнера. Колчин Степан Сергеевич ЭФМО-02-25.

## Описание проекта

Реализована контейнеризация backend-приложения на Go с помощью Docker. Сервис `tasks` упакован в Docker-образ, запускается через `docker run` и через Docker Compose.

### Эндпоинты

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET | `/health` | Проверка работоспособности сервиса |


## Структура проекта
```text
prac7-docker/
├── services/
│ └── tasks/
│ ├── cmd/
│ │ └── tasks/
│ │ └── main.go
│ ├── go.mod
│ ├── go.sum
│ ├── Dockerfile
│ └── .dockerignore
├── deploy/
│ └── docker-compose.yml
└── README.md
```

## Требования

- Go 1.23+
- Docker Desktop
- Docker Compose (встроен в Docker Desktop)

### Запуск контейнера
```bash
docker run --rm -p 8082:8082 -e TASKS_PORT=8082 techip-tasks:0.1
```
### Запуск через Docker Compose
```bash
cd deploy
docker compose up -d --build
```

## Отчётные материалы

### Запуск контейнера

<img width="974" height="77" alt="image" src="https://github.com/user-attachments/assets/f7fa30da-3487-4b7a-997e-66ebadf7ebd8" />

`curl http://localhost:8082/health`

<img width="917" height="81" alt="image" src="https://github.com/user-attachments/assets/6d950168-e8da-4ba9-b6f8-3867ec6ef686" />

### Проверка через Compose`

<img width="974" height="786" alt="image" src="https://github.com/user-attachments/assets/c698a41a-3119-44aa-9886-6c75f25811d8" />

`curl http://localhost:8082/health`

<img width="941" height="84" alt="image" src="https://github.com/user-attachments/assets/4d9ca936-4ad4-43f7-bd5e-0691ca1c9016" />


### Преверка статуса и логов

<img width="974" height="342" alt="image" src="https://github.com/user-attachments/assets/94e2f29e-fc29-4d85-91b7-1bbbf631b13c" />
