# Сервис обогащения персональных данных

## Описание проекта

Сервис предоставляет REST API для работы с персональными данными, обогащая их информацией о возрасте, поле и национальности из внешних API. Реализован на Go с использованием Gin framework.

## Технологический стек

- **Язык**: Go 1.21+
- **Фреймворк**: Gin
- **База данных**: PostgreSQL
- **Логирование**: Zap
- **Миграции**: golang-migrate
- **Контейнеризация**: Docker + Docker Compose
- **Архитектура**: DDD (Domain-Driven Design)
- **Документация**: Swagger

## Функциональные требования

### 1. REST API

#### 1.1 Получение данных
- Поддержка фильтрации по:
  - Имени/фамилии/отчеству
  - Возрасту (диапазон)
  - Полу
  - Национальности
- Пагинация (limit/offset)

```http
POST /api/users/find
```

#### 1.2 Удаление записи
```http
DELETE /api/users/delete/:user_id
```

#### 1.3 Обновление записи
```http
POST /api/users/update/:user_id
{
  "user_id": 9,
  "name": "Tania",
  "surname": "Savina",
  "patronymic": "Valerievna",
  "age": 49,
  "gender": "female",
  "country": {
      "country_id": "RU",
      "probability": 1
  }
}
```

#### 1.4 Добавление новой записи
```http
POST /api/users/create
{
  "name": "Dmitriy",
  "surname": "Ushakov",
  "patronymic": "Vasilevich" // optional
}
```

### 2. Обогащение данных

Сервис автоматически обогащает данные из внешних API:

1. **Возраст**: https://api.agify.io/?name=Dmitriy
2. **Пол**: https://api.genderize.io/?name=Dmitriy
3. **Национальность**: https://api.nationalize.io/?name=Dmitriy

Пример ответа:
```json
{
  "user_id": 9,
  "name": "Tania",
  "surname": "Savina",
  "patronymic": "Valerievna",
  "age": 49,
  "gender": "female",
  "country": {
      "country_id": "RU",
      "probability": 1
  }
}
```

### 3. База данных

- Используется PostgreSQL
- Структура БД создается через миграции
- Основная таблица `users` содержит:
  - ID
  - Имя, фамилия, отчество
  - Возраст
  - Пол
  - Даты создания
- Таблица `user_country` содержит:
  - ID
  - user_id (Является внешним ключом к таблице users)
  - Страна, к которой принадлежит пользователь
  - Вероятность принадлежности

### 4. Логирование
- **Debug-логи**: детальная информация о работе сервиса
- **Info-логи**: основные события (запросы, обновления)
- **Error-логи**: ошибки с stacktrace

### 5. Конфигурация

Все настройки вынесены в `.env` файл:
```env
REST_PORT: 8081
REST_HOST: "0.0.0.0"

POSTGRES_HOST: postgres #"127.0.0.1"
POSTGRES_PORT:  5432
POSTGRES_USER: "root"
POSTGRES_PASS: "1234"
POSTGRES_DB: "postgres"
POSTGRES_MAX_CONN: 10
POSTGRES_MIN_CONN: 5
```

### 6. Документация

Автоматически генерируемая Swagger документация доступна по:
```
GET /swagger/index.html
```

## Архитектура (DDD)

```
.
├── cmd
│   └── main.go
├── config
│   └── config.yaml
├── db
│   ├── migrations
│   │   ├── 000001_init_db.down.sql
│   │   └── 000001_init_db.up.sql
│   └── query.go
├── docker
│   └── Dockerfile
├── docker-compose.yaml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── names
│   │   ├── handler
│   │   │   └── handler.go
│   │   ├── model
│   │   │   └── model.go
│   │   ├── names.go
│   │   ├── repository
│   │   │   └── repository.go
│   │   └── service
│   │       └── service.go
│   └── transport
│       └── rest
│           ├── router.go
│           └── server.go
├── pkg
│   ├── api
│   │   └── nameData
│   │       └── utils.go
│   ├── logger
│   │   └── logger.go
│   └── postgres
│       └── postgres.go
└── README.md

21 directories, 24 files
```

## Запуск проекта

1. Собрать и запустить:
```bash
docker-compose up --build
```

2. Доступ к API:
```
http://localhost:8081/*any
```

3. Документация:
```
http://localhost:8081/swagger/index.html
```

## Дополнительные реализованные требования
- Валидация входящих данных
- Обработка ошибок
- Graceful shutdown
- Health-check endpoints