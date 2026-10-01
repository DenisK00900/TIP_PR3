# Практическое занятие №3
по дисциплине «Технологии индустриального программирования».

# Требования
- Go 1.22 или выше
- Git

# Структура проекта
```text
F:.
│   go.mod
│
├───cmd
│   └───server
│           main.go
│
└───internal
    ├───api
    │       handlers.go
    │       middleware.go
    │       responses.go
    │
    └───storage
            memory.go
```

# Скачивание и запуск
```text
git clone https://github.com/DenisK00900/TIP_PR3.git
cd TIP_PR3
go run .
```

# Запросы
```text
curl -i http://localhost:8080/
curl -i http://localhost:8080/health

curl -i http://localhost:8080/tasks //Список всех задачь

curl -i -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d "{\"title\":\"Купить молоко\"}" //Добавить задачу

curl -i "http://localhost:8080/tasks?q=молоко" //Поиск по ключевым словам

curl -i http://localhost:8080/tasks/1 //Поиск по индексу

curl -i -X PATCH http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d "{\"done\":true}" //Пометить как выполненную

curl -i -X DELETE http://localhost:8080/tasks/1 //Удаление задачи

```
(или другой порт, если он был изменён)
