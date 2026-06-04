## Практическое занятие №6 Реализация защиты от CSRF/XSS. Работа с secure cookies. Колчин Степан Сергеевич ЭФМО-02-25.

## Описание проекта

Реализовано учебное web-приложение на Go с защитой от CSRF и XSS, использованием безопасных cookies.

### Эндпоинты

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET | `/login` | Имитация входа, установка session cookie |
| GET | `/profile` | Форма редактирования профиля (с CSRF-токеном) |
| POST | `/profile` | Обновление имени с проверкой CSRF-токена |
| GET | `/hello` | Страница приветствия с безопасным выводом имени |

---

## Требования

- Go 1.16+

## Структура проекта

```text
prac6-web-security/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── cookie.go
│   │   └── csrf.go
│   ├── httpapi/
│   │   └── handler.go
│   └── store/
│       └── store.go
├── templates/
│   ├── profile.html
│   └── hello.html               
├── go.mod
└── README.md
```

## Запуск приложения

```bash
go run ./cmd/server
```

## Отчётные материалы

`http://localhost:8080/hello`

<img width="708" height="311" alt="image" src="https://github.com/user-attachments/assets/91bff3fb-29ba-4349-b9e5-142d7191cd80" />

`http://localhost:8080/profile`

<img width="633" height="333" alt="image" src="https://github.com/user-attachments/assets/6368c2fd-9336-478b-bdb5-d410c3a00bca" />

## Редактирование Имени

`http://localhost:8080/hello`

<img width="664" height="247" alt="image" src="https://github.com/user-attachments/assets/a1d31560-0bd5-4233-b80a-4b0609dca6d9" />

`http://localhost:8080/profile`

<img width="630" height="297" alt="image" src="https://github.com/user-attachments/assets/75eba9d6-3e33-4c59-bc0e-d7d444dc7fd3" />

## Ошибка CSRF

`<input type="hidden" name="csrf_token" value="asdasd">`

<img width="298" height="115" alt="image" src="https://github.com/user-attachments/assets/f57c5369-aa6d-4df5-840a-ca4a832f5b57" />


## Демонстрация XSS-безопасности

<img width="974" height="342" alt="image" src="https://github.com/user-attachments/assets/e193d719-c353-45a3-aae5-6d0d6c5d5e02" />

## опасный XSS-вариант

```go
func unsafeHello(w http.ResponseWriter, name string) {
    html := "<html><body><h1>Здравствуйте, " + name + "!</h1></body></html>"
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Write([]byte(html))
}
```

>Если name = ```<script>alert('XSS')</script>```, то скрипт выполнится в браузере
