# Практическая работа 12

Автоматическая документация Swagger.

Продолжен Notes API из практики 11. Над обработчиками добавлены аннотации swag, автоматически созданы docs.go, swagger.json и swagger.yaml. Swagger UI публикуется на /docs/index.html и использует /docs/doc.json. Генератор и библиотека swag закреплены на одной версии 1.16.6.

## Запуск на macOS или Linux

Go 1.27.1, curl. Каждая практика — отдельный Go-модуль. Команды выполняются из этой папки.

```bash
make swagger
go test ./...
go vet ./...
go run ./cmd/api
```

Файл `.env` читается оболочкой через `source`; автоматически сервер его не загружает. Порт API по умолчанию 8080. При одновременном запуске меняйте HTTP_ADDR, APP_ADDR (ПР9) или APP_PORT (ПР10).

## Проверки

Фактические результаты находятся в `evidence/`. Полные curl-команды и ответы — в `evidence/requests.txt` (для практик с HTTP-проверками). Тесты 15 и 16 записаны в `evidence/tests.txt` и `evidence/integration.txt`.

Маршруты: POST/GET `/api/v1/notes`, GET/PATCH/DELETE `/api/v1/notes/{id}`. Поиск `?q=Go`, страницы `?page=1&limit=20`. Тело создания: `{"title":"Go","content":"Текст"}`.

Swagger UI: http://localhost:8080/docs/index.html. Спецификация Swagger 2.0: docs/swagger.json и docs/swagger.yaml. Генерация code-first: `make swagger`.

## Подготовка к защите

OpenAPI описывает контракт API, Swagger UI визуализирует его и позволяет отправлять запросы. Здесь swag генерирует спецификацию Swagger 2.0.

При code-first источником служат аннотации Go. При schema-first сначала пишется контракт, затем генерируются типы и интерфейсы.

@Param описывает вход, @Success и @Failure — ответы, @Router — путь и метод.

После изменения маршрута необходимо выполнить make swagger и закоммитить docs. CI может проверить отсутствие diff после генерации.
