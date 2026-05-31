## Практическое занятие №2 gRPC: создание простого микросервиса, вызовы методов. Колчин Степан Сергеевич ЭФМО-02-25

## Описание проекта

Реализован gRPC-микросервис **StudentService** с двумя методами:

- `Ping` — проверка доступности сервера
- `GetStudentByID` — получение информации о студенте по ID

Сервер хранит данные в памяти. Клиент подключается к серверу и вызывает методы.

## Эндпоинты (gRPC методы)

| Метод | Запрос | Ответ |
|-------|--------|-------|
| Ping | `PingRequest{message}` | `PingResponse{message}` |
| GetStudentByID | `GetStudentRequest{id}` | `GetStudentResponse{student}` |

---

## Требования

- Go 1.16+


## Отчётные материалы

## Запуск сервера
```bash
go run ./cmd/server
```

<img width="845" height="89" alt="image" src="https://github.com/user-attachments/assets/8b8026ac-fbad-4c71-93ec-7d799f544333" />

## Запуск клиента
```bash
go run ./cmd/client
```

<img width="880" height="128" alt="image" src="https://github.com/user-attachments/assets/5f8f63f1-7f0d-4972-94bd-f6f67e0bbfa2" />

## Запуск клиента (Сценарий ошибки c Id: 99)

<img width="877" height="142" alt="image" src="https://github.com/user-attachments/assets/33cd1721-d34b-41fb-8dbb-153893c39ed8" />

## Структура проекта

```text
prac2-grpc/
├── proto/
│   └── student.proto
├── gen/
│   └── studentpb/
│       ├── student.pb.go
│       └── student_grpc.pb.go
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── client/
│       └── main.go
├── internal/
│   └── student/
│       ├── data.go
│       └── service.go
├── photos/                  
│   
├── go.mod
└── go.sum
