package main

import (
	"fmt"
	"os"
)

func CheckTheMainDirectory () {
	modelDir, err := os.Open("./internal/app/models")
	if err != nil {
		os.Mkdir("./internal/app/models", 0755)
	}

	modelDir.Close()
}



func main() {
	args := os.Args[1:]
	CheckTheMainDirectory()

	//TODO: реализовать считывание конфигурационного файла (варианты БД, путь до директории модели, и store)

	switch args[0] {
	case "create":
		CreateModel(args[1])

	// TODO: реализация команды добавления поля/ей в модель из командной строки/файла
	// TODO: реализация команды добавления методов в репозиторий модели с генерацией кода (например получение данных по полю из БД и возвращение их в виде массива моделей)

	default:
		fmt.Println("models create <model_name>")
	}
}
