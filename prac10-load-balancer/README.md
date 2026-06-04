## Практическое занятие №10 Горизонтальное масштабирование: использование Load Balancer (NGINX). Колчин Степан Сергеевич ЭФМО-02-25.

## Описание проекта

Реализовано горизонтальное масштабирование backend-приложения `tasks` путём запуска нескольких реплик (2 экземпляра) и распределения трафика через NGINX в роли балансировщика нагрузки.

### Архитектура

```text
Клиент → NGINX (port 8080) → tasks_1:8082 → tasks_2:8082
```

### Эндпоинты

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET | `/health` | Проверка состояния сервиса (возвращает instance ID) |
| GET | `/v1/tasks` | Получение списка задач (возвращает instance ID в заголовке) |


## Требования

- Go 1.23+
- Docker / Docker Compose
- NGINX (используется образ `nginx:1.27-alpine`)

## Структура проекта

```text
prac10-load-balancer/
├── services/
│   └── tasks/
│       ├── cmd/
│       │   └── server/
│       │       └── main.go
│       ├── go.mod
│       ├── go.sum
│       └── Dockerfile
├── deploy/
│   └── lb/
│       ├── docker-compose.yml
│       └── nginx.conf
├── photos/                  
└── README.md
```

### Сборка и запуск стенда

```bash
cd deploy/lb
docker compose up -d --build
```

## Конфигурация NGINX
Файл [deploy/lb/nginx.conf](./deploy/lb/nginx.conf):
```nginx
events {}

http {
    upstream tasks_backend {
        server tasks_1:8082;
        server tasks_2:8082;
    }

    server {
        listen 8080;

        location / {
            proxy_pass http://tasks_backend;
            proxy_set_header Host $host;
            proxy_set_header X-Request-ID $request_id;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header Authorization $http_authorization;
        }
    }
}
```

## Docker Compose
Файл [deploy/lb/docker-compose.yml](./deploy/lb/docker-compose.yml):
```yaml
version: "3.9"

services:
  tasks_1:
    build:
      context: ../../services/tasks
    container_name: tasks_1
    environment:
      APP_PORT: "8082"
      INSTANCE_ID: "tasks-1"

  tasks_2:
    build:
      context: ../../services/tasks
    container_name: tasks_2
    environment:
      APP_PORT: "8082"
      INSTANCE_ID: "tasks-2"

  nginx:
    image: nginx:1.27-alpine
    container_name: nginx_lb
    ports:
      - "8080:8080"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf:ro
    depends_on:
      - tasks_1
      - tasks_2
```


## Отчётные материалы

`Запуск стенда`

<img width="869" height="402" alt="image" src="https://github.com/user-attachments/assets/7d8e853a-0c17-4d52-a72f-6f56588a6d23" />

`Проверка health endpoint`

<img width="883" height="280" alt="image" src="https://github.com/user-attachments/assets/4d21f317-97f3-4ff6-adc7-4e803f4bed13" />

`Проверка балансировки`

<img width="886" height="888" alt="image" src="https://github.com/user-attachments/assets/e730252b-a1bb-4cd1-9028-5e8eadaa059f" />

`Проверка отказоустойчивости`

- Остановим одну реплику:

<img width="878" height="216" alt="image" src="https://github.com/user-attachments/assets/913f63cc-ba24-4446-a089-fa085de0af06" />

- Теперь X-Instance-ID только tasks-2

<img width="891" height="958" alt="image" src="https://github.com/user-attachments/assets/03f7b6d3-01e2-413e-a041-5cbb7ef84673" />

`Возврат второй реплики в работу`

<img width="880" height="206" alt="image" src="https://github.com/user-attachments/assets/adf6b9b9-32a8-4b3a-8e9e-86fa66f55a2c" />

 - Идет балансировка между двумя экземплярами

<img width="881" height="884" alt="image" src="https://github.com/user-attachments/assets/5fa0ab7a-5b34-416f-a416-8d42f9e56e03" />
