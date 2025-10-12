package main

import (
	"GOlang/algo"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Использование: calculator <число1> <операция> <число2>")
		fmt.Println("Пример: ./main.go 10 + 5")
		fmt.Println("Доступные операции: +, -, *, /, `, sqrt")
		os.Exit(1)
	}

	// Парсим аргументы
	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Ошибка: первое число '%s' невалидное\n", os.Args[1])
		os.Exit(1)
	}

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Printf("Ошибка: второе число '%s' невалидное\n", os.Args[3])
		os.Exit(1)
	}

	operation := os.Args[2]

	// Вычисляем результат
	result, err := algo.Calculator(a, b, operation)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}

	// Выводим результат
	fmt.Printf("🎯 Результат: %.2f %s %.2f = %.2f\n", a, operation, b, result)

}

//CALCULATOR!!!!
