package main

import (
	"flag"
	"fmt"
	"os"
)

// функция создания файла
func createFile(filename string, text string) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Не получилось создать:", err)
		os.Exit(1)
	}

	defer file.Close()
	writeFile(text, file)
}

// функция записи файла
func writeFile(text string, file *os.File) {
	file.WriteString(text)
}

// размер файла
func sizeFile(filename string) {
	f, err := os.Open(filename)

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File size:", fi.Size(), "bytes")
}

func main() {
	// Определяем флаги
	createFlag := flag.Bool("create", false, "Create file if it doesn't exist")
	sizeFlag := flag.Bool("size", false, "Print file size")
	textFlag := flag.String("text", "Hello, Go!", "Text to write to the file")

	flag.Parse()

	// Получаем имя файла из аргументов
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("Filename is required")
		return
	}

	if *createFlag {
		createFile(filename, *textFlag)
	}

	if *sizeFlag {
		sizeFile(filename)
	}
}
