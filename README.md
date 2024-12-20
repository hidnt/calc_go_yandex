# calc_go_yandex
## Финальное задание Yandex LMS 1 спринт.

    Данная программа вычисляет значение арифметического выражения.
    Программа поддерживает ввод рациональных чисел и арифметичексие операции (+ - * /).

## Способы запуска программы
+ Сервер. Подается POST запрос в виде JSON. `{"expression":"Выражение"}`
+ Консоль. Подается арифметическое выражение в виде строки. Чтобы использовать необходимо в файле по направлению"./cmd/main.go" изменить application.RunServer() на application.Run() 

## Чтобы запустить программу, необходимо:
1) Скачать актуальную версию `git clone git@github.com:hidnt/calc_go_yandex.git`
2) Перейти в созданную папку `cd calc_go_yandex`
3) Запустить программу `go run ./cmd/main.go`

## Чтобы запустить тесты, необходимо:
1) Скачать актуальную версию `git clone git@github.com:hidnt/calc_go_yandex.git`
2) Перейти в созданную папку `cd calc_go_yandex`
3) Запустить тестирование `go test -v ./...`

## Примеры работы программы:
    curl -X POST -H "Content-Type: application/json" -d "{ \"expression\": \"(2+2-(-2+7)*2)/2\" }" http://localhost:8080/api/v1/calculate

Возвращает
{
	"result": "-3.000000"
}
Код 200.

    curl -X POST -H "Content-Type: application/json" -d "{ \"expression\": \"123-(8*4\" }" http://localhost:8080/api/v1/calculate

Возвращает
{
	"error": "Expression is not valid"
}
Код 422.

    curl -X POST -H "Content-Type: application/json" -d "" http://localhost:8080/api/v1/calculate

Возвращает
{
	"error": "Internal server error"
} Код 500.
