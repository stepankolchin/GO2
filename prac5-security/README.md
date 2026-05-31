## Практическое занятие №5 Реализация HTTPS (TLS-сертификаты). Защита от SQL-инъекций. Колчин Степан Сергеевич ЭФМО-02-25.

## Описание проекта

Реализовано backend-приложение на Go с поддержкой HTTPS и безопасной работой с PostgreSQL через параметризованные запросы и prepared statements.

### Эндпоинты

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET | `/health` | Проверка работоспособности сервиса |
| GET | `/students?id={id}` | Получение студента по ID (безопасный) |

---

## Требования

- Go 1.16+
- Docker (для PostgreSQL)
- OpenSSL (для генерации сертификатов)

### Структура проекта
```text
prac5-security/
├── cmd/
│   └── server/
│       └── main.go
├── certs/
│   ├── server.crt
│   └── server.key
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── httpapi/
│   │   └── handler.go
│   └── student/
│       ├── model.go
│       └── repo.go
├── photos/                  
├── go.mod
├── go.sum
└── README.md
```

### Запуск приложения

```bash
go run ./cmd/server
```

## Отчётные материалы

`GET /health`

<img width="455" height="86" alt="image" src="https://github.com/user-attachments/assets/b935a187-86cb-4264-bdc0-e185ae3f8fd9" />

`GET /students?id=1`

<img width="878" height="102" alt="image" src="https://github.com/user-attachments/assets/dc0d5d31-5364-4af4-a1e7-2f5b2420a1ba" />

`GET /students?id=1000(student not found)`

<img width="570" height="88" alt="image" src="https://github.com/user-attachments/assets/72c41cf0-a389-466c-b95a-c7cd808f030b" />


### Безопасная работа с SQL

`Опасный подход`

```go
// Конкатенация строк с пользовательским вводом
query := "SELECT * FROM students WHERE id = " + rawID
// При rawID = "1 OR 1=1" запрос превращается в:
// SELECT * FROM students WHERE id = 1 OR 1=1
// Возвращает ВСЕХ студентов!
```

`Безопасный подход`

```go
// Значение передаётся отдельно от SQL
row := db.QueryRow(
    "SELECT id, full_name, study_group, email FROM students WHERE id = $1",
    id,
)
```

`Prepared Statement`

```go
stmt, _ := db.Prepare("SELECT * FROM students WHERE id = $1")
defer stmt.Close()
row := stmt.QueryRow(id)  // многократное использование
```
