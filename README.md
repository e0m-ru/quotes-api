# Quote REST API Service
## Мини-сервис “Цитатник”
Простой REST API сервис для управления цитатами, написанный на Go.

## API Endpoints

- `POST /quotes` - Добавить новую цитату
- `GET /quotes` - Получить все цитаты (можно фильтровать по автору)
- `GET /quotes/random` - Получить случайную цитату
- `DELETE /quotes/{id}` - Удалить цитату по ID

## Примеры использования

Добавить цитату:
```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

Получить все цитаты:
```bash
curl http://localhost:8080/quotes
```

Получить случайную цитату:
```bash
curl http://localhost:8080/quotes/random
```

Фильтрация по автору:
```bash
curl http://localhost:8080/quotes?author=Confucius
```

Удалить цитату:
```bash
curl -X DELETE http://localhost:8080/quotes/1
```

## Требования

- Go 1.16 или выше

## Установка и запуск
1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/e0m-ru/quotes-api.git
   cd quote-api
2. Установите зависимости:
   ```bash
   go mod download
   ```
3. Запустите сервер:
   ```bash
   go run .
   ```
Сервер будет доступен по адресу `http://localhost:8080`.

## Тестирование
Для запуска тестов выполните:
```bash
go test -v
```

## Лицензия
MIT
