## Практическое занятие №11 Создание GraphQL API с использованием gqlgen. Запросы и мутации. Колчин Степан Сергеевич ЭФМО-02-25.

## Описание проекта

Реализован GraphQL API для сущности `Task` с использованием библиотеки `gqlgen`. API поддерживает запросы на чтение (получение списка и одной задачи) и мутации (создание, обновление, удаление задач). Данные хранятся in-memory.

## GraphQL-схема

```graphql
type Task {
    id: ID!
    title: String!
    description: String
    done: Boolean!
}

type Query {
    tasks: [Task!]!
    task(id: ID!): Task
}

type Mutation {
    createTask(input: CreateTaskInput!): Task!
    updateTask(id: ID!, input: UpdateTaskInput!): Task!
    deleteTask(id: ID!): Boolean!
}

input CreateTaskInput {
    title: String!
    description: String
}

input UpdateTaskInput {
    title: String
    description: String
    done: Boolean
}
```

## Требования
- Go 1.23+
- gqlgen v0.17.70

## Структура проекта

```text
prac11-graphql/
├── graph/
│   ├── model/
│   │   └── models_gen.go
│   ├── resolver.go
│   ├── schema.resolvers.go
│   ├── schema.graphqls
│   ├── generated.go
│   └── store.go
├── photos/ 
├── server.go
├── gqlgen.yml
├── go.mod
├── go.sum
└── README.md
```

## Запуск сервера
```bash
go run server.go
```


### Доступные операции

| Тип | Операция | Описание |
|-----|----------|----------|
| Query | `tasks` | Получение списка всех задач |
| Query | `task(id)` | Получение задачи по ID |
| Mutation | `createTask(input)` | Создание новой задачи |
| Mutation | `updateTask(id, input)` | Обновление задачи |
| Mutation | `deleteTask(id)` | Удаление задачи |



## Отчётные материалы

`Запрос списка задач`

<img width="974" height="427" alt="image" src="https://github.com/user-attachments/assets/87013936-1040-4ca5-bc2f-8893a67b2759" />

`получение задачи по ID`

<img width="915" height="595" alt="image" src="https://github.com/user-attachments/assets/97b9df56-d2ff-424a-af49-6c8fdaad11c5" />


`Создание новой задачи (мутация)`

<img width="974" height="636" alt="image" src="https://github.com/user-attachments/assets/0ecc91d1-0549-4a7f-94f7-832310bf0500" />


`Обновление задачи`

<img width="974" height="659" alt="image" src="https://github.com/user-attachments/assets/b6856f07-ec0c-4a1d-ad21-3b71327be5cb" />

`Удаление задачи`

<img width="974" height="748" alt="image" src="https://github.com/user-attachments/assets/43ca74c9-8c6c-4ff2-8bf5-e19206fb133b" />
