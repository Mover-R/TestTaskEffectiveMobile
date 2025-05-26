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
- Сортировка по любому полю

```http
GET /api/v1/persons?age_from=20&age_to=30&gender=male&limit=10&offset=0
```

#### 1.2 Удаление записи
```http
DELETE /api/v1/persons/{id}
```

#### 1.3 Обновление записи
```http
PATCH /api/v1/persons/{id}
Content-Type: application/json

{
  "name": "UpdatedName",
  "age": 35
}
```

#### 1.4 Добавление новой записи
```http
POST /api/v1/persons
Content-Type: application/json

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
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "Dmitriy",
  "surname": "Ushakov",
  "patronymic": "Vasilevich",
  "age": 34,
  "gender": "male",
  "nationality": "RU",
  "created_at": "2023-05-26T12:00:00Z",
  "updated_at": "2023-05-26T12:00:00Z"
}
```

### 3. База данных

- Используется PostgreSQL
- Структура БД создается через миграции
- Основная таблица `persons` содержит:
  - ID (UUID)
  - Имя, фамилия, отчество
  - Возраст
  - Пол
  - Национальность
  - Даты создания/обновления

### 4. Логирование

- **Debug-логи**: детальная информация о работе сервиса
- **Info-логи**: основные события (запросы, обновления)
- **Error-логи**: ошибки с stacktrace

### 5. Конфигурация

Все настройки вынесены в `.env` файл:
```
POSTGRES_HOST=db
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=persons
SERVER_PORT=8080
ENVIRONMENT=development
```

### 6. Документация

Автоматически генерируемая Swagger документация доступна по:
```
GET /swagger/index.html
```

## Архитектура (DDD)

```
.
├── cmd/              # Точка входа
├── internal/
│   ├── config/       # Конфигурация
│   ├── domain/       # Доменные модели
│   ├── application/  # Use cases
│   ├── infrastructure/
│   │   ├── api/      # Внешние API клиенты
│   │   ├── db/       # Репозитории
│   │   └── http/     # HTTP handlers
│   └── interfaces/   # REST API
├── pkg/
│   ├── logger/       # Логирование
│   └── postgres/     # DB подключение
├── db/
│   └── migrations/   # SQL миграции
├── docs/             # Swagger docs
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

## Запуск проекта

1. Собрать и запустить:
```bash
docker-compose up --build
```

2. Применить миграции:
```bash
docker-compose run --rm migrate up
```

3. Доступ к API:
```
http://localhost:8080/api/v1/persons
```

4. Документация:
```
http://localhost:8080/swagger/index.html
```

## Дополнительные требования

- Валидация входящих данных
- Обработка ошибок
- Тесты (unit, integration)
- Graceful shutdown
- Health-check endpoints
- Rate limiting для внешних API