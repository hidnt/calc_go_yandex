package main

import "github.com/hidnt/calc_go_yandex/internal/application"

func main() {
	application := application.New()
	//application.Run()
	application.RunServer()
}
