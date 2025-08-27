# webhook-bypass

Простой сервис для пересылки POST-запросов на указанный webhook (например, Discord).

## Запуск

1. Установите зависимости:
   ```bash
   go mod tidy
   ```

2. Запустите сервер:
   ```bash
   go run ./cmd/main.go
   ```
   Сервер будет слушать порт `9090`.

## Пример использования

Отправьте POST-запрос на сервер, указав URL webhook в параметре `webhook`:

```bash
curl -X POST "http://localhost:9090/?webhook=https://discord.com/api/webhooks/ВАШ_WEBHOOK" \
-H "Content-Type: application/json" \
-d '{"content": "Hello from webhook-bypass!"}'
```

**Для Discord:**  
Поле `content` обязательно, иначе сообщение не будет отправлено.

## Описание

- Сервер принимает POST-запрос с JSON-телом.
- Пересылает тело запроса на указанный webhook методом POST.
- Возвращает ответ от webhook.

## Лицензия