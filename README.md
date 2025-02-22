# Module-1,2
Итоговая задача модуля 1+2 Яндекс лицея

Этот проект реализует веб-сервис, принимающий выражение через Http запрос и возвращабщий результат вычислений

Инструкция по запуску:

Убедитесь, что у вас установлен Go (версия 1.16 или выше).

Скопируйте репозиторий(через git bash):

```bash
git clone https://github.com/Killered672/Module1
```

```bash
cd Module1
```

Запустите сервер:

```bash
go run ./cmd/calc_service/main.go
```

Сервер будет доступен по адресу http://localhost:8080.

У меня чтобы дальше работали запросы нужно перезапустить консоль git bash, затем опять открыть путь вводом cd Module1 и после этого можно вводить запросы

Примеры использования:

Успешный запрос:

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2*2+2"
}'
```

Ответ:

```bash
{
  "id": "..."
}
```
После можно посмотреть этап выполнения данного запроса и его результат(если уже вычислилось):

```bash
curl --location 'http://localhost:8080/api/v1/expressions'
```
Вывод:
```bash
{"expressions":[{"id":"1740240110508066400","status":"pending"}]}
```
Или узнать точный результат нужного выражения по его точному id:

```bash
curl --location 'http://localhost:8080/api/v1/expressions/:id'
```

Ошибка 422 (невалидное выражение):

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2+a"
}'
```
Ответ:

```bash
{
  "error": "Expression is not valid"
}
```

Ошибка 500 (внутренняя ошибка сервера):

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2/0"
}'
```
Ответ:

```bash
{
  "error": "Internal server error"
}
```

Тесты для evaluator запускаются тоже через git bash(или можно через visual studio code):

```bash
go test ./internal/evaluator
```

при успешном прохождение теста должен вывестись ответ:

```bash
ok  	calc_service/internal/evaluator	0.001s
```

при ошибке в тестах будет указано где она совершена.
