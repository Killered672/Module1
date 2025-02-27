# Module-1,2
Итоговая задача модуля 1+2 Яндекс лицея.

Этот проект реализует веб-сервис, принимающий выражение через Http запрос и возвращабщий результат вычислений.

Инструкция по запуску:

1)Убедитесь, что у вас установлен Go (версия 1.16 или выше).

2)Скопируйте репозиторий(через git bash ):

```bash
git clone https://github.com/Killered672/Module2calc
```

```bash
cd Module2calc
```

Запускаем orchestator:

```bash
export TIME_ADDITION_MS=200
export TIME_SUBTRACTION_MS=200
export TIME_MULTIPLICATIONS_MS=300
export TIME_DIVISIONS_MS=400

go run cmd/orchestrator.start/main.go
```

Вы получите ответ  Starting Orchestrator on port 8080.

В новом bash(у меня так,может у вас будет дотупно и в одном и том же ):

Опять переходим в репозиторию с проектом:

```bash
cd Module2calc
```

Затем запускаем agent:

```bash
export COMPUTING_POWER=4
export ORCHESTRATOR_URL=http://localhost:8080

 go run cmd/agent.start/main.go
```

Вы получите ответ:
Starting Agent...
Starting worker 0
Starting worker 1
Starting worker 2
Starting worker 3

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

После можно посмотреть этап выполнения данного запроса и его результат(если уже вычислилось ):

```bash
curl --location 'http://localhost:8080/api/v1/expressions'
```

Вывод:

```bash
{"expressions":[{"id":"1740240110508066400","expression":"2*2+2,"status":"pending"}]}
```

Если вычисления выполнены то:

```bash
{"expression":{"id":"1","expression":"2*2+2","status":"completed","result":6}}
```

Или узнать точный результат нужного выражения по его точному id:

```bash
curl --location 'http://localhost:8080/api/v1/expressions/id'
```

Ошибки при запросах:

Ошибка 404(отсутствие выражения ):

```bash
{"error":"Expression not found"}
```

Ошибка 422 (невалидное выражение ):

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
  {"error":"expected number at position 2"}
}
```

Ошибка 500 (внутренняя ошибка сервера ):

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2/0"
}'
```
Ответ(у  меня высвечивается изначально id созданной задачи,а после в git bash где был запущен agent.start можно увидеть что выводится деление на 0 ):

```bash
{
  Worker n: error computing task 3: division by zero
}
```

Тесты для agent запускаются тоже через git bash(или можно через visual studio code):

1)Сначала опять переходим в папку с модулем.

```bash
cd Module2calc
```

2)Затем запускаем тестирование:

```bash
go test ./internal/agent/agent_calculation_test.go
```

3)При успешном прохождение теста должен вывестись ответ:

```bash
ok  	calc_service/internal/evaluator	0.001s
```

4)При ошибке в тестах будет указано где она совершена.
P.S ошибка связанная с не указанным ErrDivivsionByZero появляется так как в функции тестирования я ее не оглашаю,
она создает конфликты в visual studio code так как уже присутствует в самом агенте

Мой тг для связи: @Killered_656(можно писать не только по проекту)