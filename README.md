# Module-1
Итоговая задача модуля 1 Яндекс лицея 

Инструкция по запуску
Убедитесь, что у вас установлен Go (версия 1.16 или выше).

Склонируйте репозиторий(через git bash):

git clone https://github.com/yourusername/calc_service.git
cd calc_service

Запустите сервер:

go run ./cmd/calc_service/main.go
Сервер будет доступен по адресу http://localhost:8080.

Примеры использования
Успешный запрос:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
Ответ:

{
  "result": "6.00"
}



Ошибка 422 (невалидное выражение):
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*a"
}'

Ответ:
json
Copy
{
  "error": "Expression is not valid"
}



Ошибка 500 (внутренняя ошибка сервера):
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2/0"
}'

Ответ:
{
  "error": "Internal server error"
}