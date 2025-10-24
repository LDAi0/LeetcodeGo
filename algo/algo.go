package algo

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	// Приводим к нижнему регистру и убираем мусор
	var cleaned strings.Builder
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleaned.WriteRune(unicode.ToLower(ch))
		}
	}

	// Два указателя
	str := cleaned.String()
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func FizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}

	}
	slices.Reverse(result)
	return result
}

func Sum(a, b int) int {
	return a + b
}

func Calculator(a float64, b float64, operation string) (float64, error) {
	switch operation {
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Деление на ноль")
		}
		return a / b, nil
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "`":
		return math.Pow(a, b), nil
	case "sqrt":
		if a < 0 {
			return 0, fmt.Errorf("корень из отрицательного числа")
		}
		return math.Sqrt(a), nil
	default:
		return 0, fmt.Errorf("Unknown operation")
	}

}

func CofeCounter(n, d int) int {
	var freeCoups int
	freeCoups = d / 7
	d -= freeCoups
	freeCoups += freeCoups / n

	return freeCoups
}

func InteractiveMode() {
	fmt.Println("🧮 Добро пожаловать в калькулятор!")
	fmt.Println("Введите 'exit' для выхода")

	for {
		fmt.Print("Введите выражение (например: 5 + 3): ")

		var a, b float64
		var op string

		_, err := fmt.Scanf("%f %s %f", &a, &op, &b)
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте снова.")
			continue
		}

		result, err := Calculator(a, b, op)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}

		fmt.Printf("Результат: %.2f\n\n", result)
	}
}
