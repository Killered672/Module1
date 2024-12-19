# Module-1
Итоговая задача модуля 1 Яндекс лицея 

Этот проект реализует веб-сервис, принимающий выражение через Http запрос и возвращабщий результат вычислений

Инструкция по запуску
Убедитесь, что у вас установлен Go (версия 1.16 или выше).

Скопируйте репозиторий(через git bash):

git clone https://github.com/Killered672/Module1

cd calc_service

Запустите сервер:

go run ./cmd/calc_service/main.go

Сервер будет доступен по адресу http://localhost:8080.

Примеры использования
Успешный запрос:

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '

{
  "expression": "2+2*2"
}'

Ответ:

{
  "result": "6.00"
}



Ошибка 422 (невалидное выражение):

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '

{
  "expression": "2+2*a"
}'

Ответ:
{
  "error": "Expression is not valid"
}



Ошибка 500 (внутренняя ошибка сервера):

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '

{
  "expression": "2/0"
}'

Ответ:
{
  "error": "Internal server error"
}