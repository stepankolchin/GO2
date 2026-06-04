# Практическое занятие №15 Деплой приложения на VPS. Настройка systemd. Колчин Степан Сергеевич ЭФМО-02-25.

## Процесс выполнения

Для выполнения практической работы использовалась подсистема WSL (Windows Subsystem for Linux) с установленным дистрибутивом Ubuntu. WSL предоставляет полноценное Linux-окружение, включая поддержку systemd, что позволяет полностью воспроизвести все этапы деплоя приложения на VPS. Все команды, приведённые в отчёте, были выполнены в терминале WSL и дают результат, идентичный работе на удалённом сервере.

### Структура размещения на сервере
```text
/opt/tasks/
└── tasks                 # исполняемый файл

/etc/tasks/
└── tasks.env             # конфигурация

/etc/systemd/system/
└── tasks.service         # unit-файл systemd
```

` Создание системного пользователя`

<img width="974" height="60" alt="image" src="https://github.com/user-attachments/assets/b3aae68d-60bf-4352-a51d-650a1a17ee51" />

`Создание директорий`

<img width="870" height="156" alt="image" src="https://github.com/user-attachments/assets/1bbee4c6-5fae-4180-90a5-4f36e1d304b4" />

`Собрали бинарник внутри WSL`

<img width="974" height="31" alt="image" src="https://github.com/user-attachments/assets/d016e8c1-d3c8-4b90-83cf-55e9c36aff2e" />

`Копирование бинарника в целевую директорию`

<img width="974" height="121" alt="image" src="https://github.com/user-attachments/assets/47ab472e-81b4-4e77-a486-8b8fe4de6f80" />

`Создание unit-файла`

<img width="974" height="327" alt="image" src="https://github.com/user-attachments/assets/e1e29220-d2b0-4f6f-831c-d6824a52168f" />

`Запуск сервиса`

<img width="974" height="372" alt="image" src="https://github.com/user-attachments/assets/40cf7971-cece-444f-b36f-ef1361a25b24" />

`Проверка доступности /v1/tasks/1`

<img width="974" height="118" alt="image" src="https://github.com/user-attachments/assets/dc50b2ff-08a5-4ad0-85e9-b22dc87c28e3" />

`Просмотр логов`

<img width="974" height="453" alt="image" src="https://github.com/user-attachments/assets/87afce84-3e73-4dbe-bc44-d16dabd63af8" />
