package main

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	filename := "testfile.txt"
	text := "Test content" // Теперь переменная используется

	// Удаляем файл перед тестом (если существует)
	os.Remove(filename)

	// Создаем файл
	if err := createFile(filename, text); err != nil { // Используем переменную text
		t.Fatalf("createFile failed: %v", err)
	}

	// Проверяем существование файла
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("File was not created: %s", filename)
	}

	// Удаляем файл после теста
	os.Remove(filename)
}
