## Практическое занятие №9 Реализация распределённого кэша (Redis cluster). Колчин Степан Сергеевич ЭФМО-02-25.

## Описание проекта

Реализован backend-сервис на Go с кэшированием данных в Redis по стратегии **cache-aside**. Сервис поддерживает CRUD-операции над задачами (tasks) с автоматической инвалидацией кэша при изменении данных и деградацией при недоступности Redis.

### Эндпоинты

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET | `/v1/tasks/{id}` | Получение задачи по ID (с кэшированием) |
| PATCH | `/v1/tasks/{id}` | Обновление задачи (с инвалидацией кэша) |
| DELETE | `/v1/tasks/{id}` | Удаление задачи (с инвалидацией кэша) |


## Требования

- Go 1.23+
- Docker / Docker Compose
- Redis 7.4

### Запуск Redis

```bash
cd deploy/redis
docker compose up -d
```

### Запуск сервера
```bash
go run ./cmd/server
```

### Структура проекта
```text
pz9-redis-cache/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── cache/
│   │   ├── keys.go
│   │   ├── redis.go
│   │   └── ttl.go
│   ├── config/
│   │   └── config.go
│   ├── httpapi/
│   │   └── handler.go
│   ├── service/
│   │   └── task_service.go
│   └── task/
│       ├── model.go
│       └── repo.go
├── deploy/
│   └── redis/
│       └── docker-compose.yml                
├── go.mod
├── go.sum
└── README.md
```

## Отчётные материалы

`Запуск Redis через Docker Compose`

<img width="974" height="375" alt="image" src="https://github.com/user-attachments/assets/88b60184-efe6-492c-9f5f-e1d9f1db601d" />

`Проверка cache miss и cache hit`

<img width="974" height="89" alt="image" src="https://github.com/user-attachments/assets/d841aec5-ca47-4965-b498-0cd087789fe0" />

`логи сервера`

<img width="891" height="102" alt="image" src="https://github.com/user-attachments/assets/a10a78b4-6aee-4831-ae15-7cde5e2c2e0b" />

`2-ой запрос`

<img width="974" height="89" alt="image" src="https://github.com/user-attachments/assets/1b9855b1-a9ba-4684-8020-2f0dae7774f3" />

<img width="495" height="106" alt="image" src="https://github.com/user-attachments/assets/8950f182-1ed3-4ee0-90d3-351df26973a5" />


`Проверка инвалидации при обновлении`

<img width="872" height="309" alt="image" src="https://github.com/user-attachments/assets/f8886b70-f79f-479c-9335-43c715c2a162" />


<img width="894" height="161" alt="image" src="https://github.com/user-attachments/assets/9cebc672-aab8-4d7a-bdfe-76970eb25d20" />

`Проверка удаления`

<img width="869" height="159" alt="image" src="https://github.com/user-attachments/assets/959e1511-33e0-41f0-8622-b9fc262d90ea" />

`Проверка деградации при остановке Redis`

<img width="880" height="209" alt="image" src="https://github.com/user-attachments/assets/fb828d09-92fa-4f10-aa42-581ae4051a4b" />

`Сервер НЕ падает, ответ приходит из репозитория (через БД)`

<img width="874" height="128" alt="image" src="https://github.com/user-attachments/assets/b0d861ce-add7-4330-98ff-0657b060d441" />

`В логах ошибка Redis, но ответ 200`

<img width="974" height="344" alt="image" src="https://github.com/user-attachments/assets/aa3083fc-40d8-4f40-a638-c98303e1fac5" />
