## Практическое занятие №3 Логирование с помощью zap. Ведение структурированных логов. Колчин Степан Сергеевич ЭФМО-02-25.

## Описание проекта

Реализовано HTTP-приложение на Go с использованием библиотеки `zap` для структурированного логирования.

### Эндпоинты

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET | `/health` | Проверка работоспособности сервиса |
| GET | `/students/{id}` | Получение информации о студенте по ID |

## Требования

- Go 1.16+

## Структура проекта
```text
prac3-logging/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── httpapi/
│   │   ├── handler.go
│   │   ├── middleware.go
│   │   └── response_writer.go
│   └── student/
│       ├── model.go
│       └── repo.go
├── pkg/
│   └── logger/
│       └── logger.go       
├── go.mod
├── go.sum
└── README.md
```
## Запуск сервера
```bash
go run ./cmd/server
```

## Отчётные материалы

`GET /health`

<img width="836" height="89" alt="image" src="https://github.com/user-attachments/assets/030cd38d-0ae6-42d5-a664-a1f7fb8fd194" />

`Логи сервера:`

<img width="974" height="99" alt="image" src="https://github.com/user-attachments/assets/a70de41c-4de3-47d7-a0ef-905e2087a5a3" />

`GET /students/1`

<img width="894" height="102" alt="image" src="https://github.com/user-attachments/assets/f8eb272a-3757-492d-a60a-8a339dcfa144" />


`Логи сервера:`

<img width="974" height="56" alt="image" src="https://github.com/user-attachments/assets/3225d996-67ea-4a86-8a29-54a9cb8b347c" />


`GET /students/a (invalid student id)`

<img width="889" height="81" alt="image" src="https://github.com/user-attachments/assets/c3c0fb36-43fa-400d-af58-5d266302ebc5" />


`Логи сервера:`

<img width="974" height="56" alt="image" src="https://github.com/user-attachments/assets/c31d9d57-dcce-4eba-af77-829349830a0e" />


`GET /students/1000 (student not found)`

<img width="864" height="84" alt="image" src="https://github.com/user-attachments/assets/8d6e2f1d-0d56-4d74-aaf7-9e73fd992edf" />


`Логи сервера:`

<img width="974" height="108" alt="image" src="https://github.com/user-attachments/assets/a94d8d11-f668-43da-aecc-7ea443b90469" />
