package main

import (
	"flag"
	"fmt"
	"os"
)

// Функция создания файла (теперь возвращает error)
func createFile(filename string, text string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create error: %w", err)
	}
	defer file.Close()

	if err := writeFile(text, file); err != nil {
		return fmt.Errorf("write error: %w", err)
	}
	return nil // Явный возврат при успехе
}

// Функция записи файла (возвращает error)
func writeFile(text string, file *os.File) error {
	_, err := file.WriteString(text)
	return err // Возвращаем ошибку или nil
}

// Функция получения размера файла
func sizeFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("open error: %w", err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return fmt.Errorf("stat error: %w", err)
	}

	fmt.Println("File size:", fi.Size(), "bytes")
	return nil
}

func main() {
	createFlag := flag.Bool("create", false, "Create file")
	sizeFlag := flag.Bool("size", false, "Print file size")
	textFlag := flag.String("text", "Hello, Go!", "Text to write")

	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("Filename is required")
		os.Exit(1)
	}

	if *createFlag {
		if err := createFile(filename, *textFlag); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	if *sizeFlag {
		if err := sizeFile(filename); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}
