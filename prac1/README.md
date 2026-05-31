## Практическое занятие №1. Разделение монолита на 2 микросервиса. Взаимодействие через HTTP. Колчин Степан Сергеевич ЭФМО-02-25.

## Описание проекта

Реализована распределённая система из двух микросервисов:

- **user-service** - хранит и отдаёт информацию о пользователях.
- **order-service** - хранит заказы, при необходимости обращается к user-service для получения данных пользователя и возвращает агрегированный ответ.

### Эндпоинты

#### user-service (порт 8081)

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET | `/users/{id}` | Получить данные пользователя по ID |

#### order-service (порт 8082)

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET | `/orders/{id}` | Получить заказ по ID (без данных пользователя) |
| GET | `/orders/{id}/full` | Получить заказ с данными пользователя (агрегированный ответ) |

## Требования

- Go 1.16+
- Git
- Два свободных порта: 8081 и 8082

## Отчётные материалы

## user-service
```bash
curl http://localhost:8081/users/1
```

<img width="565" height="70" alt="image" src="https://github.com/user-attachments/assets/4f1f1cb7-e969-45d6-a79f-2d152958cc22" />

```bash
curl http://localhost:8081/users/999
```

<img width="561" height="55" alt="image" src="https://github.com/user-attachments/assets/bb02727a-48f0-4974-9309-88e02aa680cb" />


## order-service 

```bash
curl http://localhost:8082/orders/101
```

<img width="562" height="68" alt="image" src="https://github.com/user-attachments/assets/1dd67427-058a-4e9f-bfa9-a4e7e5ba01d4" />


```bash
curl http://localhost:8082/orders/101/full
```

<img width="561" height="69" alt="image" src="https://github.com/user-attachments/assets/bef62213-a13f-49cd-b170-8d66bc72a4b8" />

## Сценарий ошибки (user-service недоступен)

```bash
curl http://localhost:8082/orders/101/full
```

<img width="559" height="100" alt="image" src="https://github.com/user-attachments/assets/4ea45686-446a-49db-ab3f-6ec98a94f9dc" />


## Структура проекта

```text
pz1-microservices/
├── user-service/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   └── user/
│   │       ├── model.go
│   │       ├── repo.go
│   │       └── handler.go
│   └── go.mod
├── order-service/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   └── order/
│   │       ├── model.go
│   │       ├── repo.go
│   │       ├── client.go
│   │       └── handler.go
│   └── go.mod                  
│   
└── README.md
```
