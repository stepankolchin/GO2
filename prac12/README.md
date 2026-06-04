## Практическое занятие №12 Сравнение REST и GraphQL: разработка одного и того же функционала двумя способами. Колчин Степан Сергеевич ЭФМО-02-25.

## REST-сервис tasks - [prac9-redis-cache](../prac9-redis-cache/)

## GraphQL-сервис - [prac11-graphql](../prac11-graphql/)

## Реализация REST API (pz9-redis-cache)

- Базовый URL: `http://localhost:8082`

`Получение задачи по ID`

<img width="974" height="89" alt="image" src="https://github.com/user-attachments/assets/991c3e92-b654-472e-bd1a-ec831fd10adf" />

>**REST** возвращает все поля, включая description и due_date, даже если они не нужны клиенту.

`Создание задачи`

<img width="560" height="65" alt="image" src="https://github.com/user-attachments/assets/b01badc2-c438-4d5d-a393-5441d2f3efaf" />

`Ошибка: задача не найдена`

<img width="869" height="159" alt="image" src="https://github.com/user-attachments/assets/45a3ab42-e0d0-4eb2-a565-7c7ad672c2f9" />

>Ответ: HTTP 404 Not Found с сообщением task not found

## Реализация GraphQL API (pz11-graphql)

`Список задач (только нужные поля)`

<img width="974" height="427" alt="image" src="https://github.com/user-attachments/assets/d0c0acdc-568f-4581-b456-8b76ead1d048" />


>GraphQL возвращает только запрошенные поля

`Детали задачи`

<img width="915" height="595" alt="image" src="https://github.com/user-attachments/assets/5d0bd2b3-465d-437e-811a-79f7f82cdb7b" />


`Создание задачи`

<img width="974" height="636" alt="image" src="https://github.com/user-attachments/assets/7066d25b-9ac5-41cd-bf13-15c6a983c231" />


`Ошибка: задача не найдена`

<img width="1191" height="447" alt="image" src="https://github.com/user-attachments/assets/41df25a2-1388-4028-8e04-586d2338c09b" />

>Ответ: HTTP 200 OK с полем data.task: null (без HTTP-ошибки)

## Сравнение по критериям

### 1. Количество запросов

| Действие | REST | GraphQL |
|----------|------|---------|
| Получить список задач | 1 запрос | 1 запрос |
| Получить детали задачи | 1 запрос | 1 запрос |
| Создать задачу | 1 запрос | 1 запрос |

### 2. Объём данных (over-fetching)

| Сценарий | REST | GraphQL |
|----------|------|---------|
| Список задач | Возвращает все поля (id, title, description, due_date) | Возвращает только запрошенные (id, title, done) |
| Наличие лишних полей |  есть (description, due_date) |  нет |

### 3. Обработка ошибок

| Ситуация | REST | GraphQL |
|----------|------|---------|
| Задача не найдена | HTTP 404 + тело ошибки | HTTP 200 + data.task: null |
| Неверный ID | HTTP 400 Bad Request | HTTP 200 + поле errors |
| Способ определения ошибки | По HTTP-статусу | По наличию поля errors или null |

### 4. Кэширование

| Аспект | REST | GraphQL |
|--------|------|---------|
| Кэширование по URL |  Просто |  Единый endpoint |
| Стандартные HTTP-средства |  Поддерживаются |  Требуются доп. решения |
| Клиентский кэш |  Не предусмотрен |  Возможен (Apollo, Relay) |

## Вывод:
В рамках работы были реализованы идентичные возможности для сущности Task с использованием двух подходов: REST и GraphQL.

REST оказался проще при разработке и отладке, а также предоставляет более удобные механизмы для кэширования.
GraphQL, в свою очередь, выигрывает за счёт гибкости запросов: клиент может запросить ровно те поля, которые нужны, — это особенно ценно для мобильных приложений и насыщенных интерфейсов.

Тем не менее, для типового CRUD-сервиса с ограниченным числом экранов REST остаётся более прагматичным выбором.
Итоговое решение зависит от конкретной задачи: для публичных API с несложной структурой лучше подходит REST, а для приложений, где на разных экранах требуются разные наборы данных, предпочтительнее GraphQL.
