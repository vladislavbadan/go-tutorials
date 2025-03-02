package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Введи: go run file_manager.go <ИМЯ ФАЙЛА>")
		return
	}

	filename := os.Args[1]

	// открываем файл
	f, err := os.Open(fmt.Sprintf("%s.txt", filename))
	// проверка на ошибка
	if os.IsNotExist(err) {
		// создаем файл так как его нет
		file, err := os.Create(fmt.Sprintf("%s.txt", filename))

		if err != nil {
			fmt.Println("Не получилось создать:", err)
			os.Exit(1)
		}
		defer file.Close()
		file.WriteString("Hello go!")

	} else {
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			panic(err)
		}

		fmt.Println("File size:", fi.Size(), "bytes")
	}
}
