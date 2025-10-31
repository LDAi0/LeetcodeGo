package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Открываем файл
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Гарантируем закрытие файла при выходе из функции
	defer file.Close()

	// Создаем сканер
	scanner := bufio.NewScanner(file)

	// Читаем файл построчно
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Проверяем ошибки при сканировании
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
