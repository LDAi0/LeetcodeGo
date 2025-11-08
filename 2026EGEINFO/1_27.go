package EGEINFO

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

// min_int64 - вспомогательная функция для поиска минимума из двух int64
func min_int64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// solveFile обрабатывает один файл и возвращает минимальную сумму
func SolveFile(filename string) int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Ошибка: Не удалось открыть файл %s: %v", filename, err)
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// 1. Считываем K (минимальное расстояние)
	if !scanner.Scan() {
		log.Println("Ошибка: Файл пуст (не удалось считать K)")
		return -1
	}
	k, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Printf("Ошибка: Не удалось прочитать K: %v", err)
		return -1
	}

	// 2. Считываем N (количество чисел)
	if !scanner.Scan() {
		log.Println("Ошибка: Файл пуст (не удалось считать N)")
		return -1
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Printf("Ошибка: Не удалось прочитать N: %v", err)
		return -1
	}
	if n == 0 {
		return 0 // Нет данных
	}

	// Используем math.MaxInt64 как "бесконечность" для поиска минимума
	const infinity = math.MaxInt64

	// Создаем "буферы" (очереди) размером K.
	// Они будут хранить минимальные значения, которые мы видели K шагов назад.

	// min1_buffer[i] будет хранить лучший (минимальный) *первый* элемент,
	// который был доступен K шагов назад.
	min1_buffer := make([]int64, k)

	// min2_buffer[i] будет хранить лучшую (минимальную) *пару* (сумма двух),
	// которая была доступна K шагов назад.
	min2_buffer := make([]int64, k)

	// Инициализируем буферы "бесконечностью"
	for i := range min1_buffer {
		min1_buffer[i] = infinity
		min2_buffer[i] = infinity
	}

	// global_min1: Минимальное *одно* число, которое мы видели (для обновления min1_buffer)
	var global_min1 int64 = infinity
	// global_min2: Минимальная *сумма двух* чисел (a[i] + a[j]), где j - i >= K
	var global_min2 int64 = infinity
	// global_min3: Минимальная *сумма трех* чисел (a[i] + a[j] + a[l]),
	// где j - i >= K и l - j >= K. Это наш итоговый ответ.
	var global_min3 int64 = infinity

	// 3. Читаем N строк с данными
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			log.Println("Ошибка: Файл закончился раньше, чем ожидалось (N)")
			break
		}

		val, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Printf("Ошибка: не удалось прочитать число на строке %d: %v", i+3, err)
			continue
		}

		// ptr - это индекс в наших круговых буферах (очередях).
		// Он указывает на ячейку, хранящую значение K шагов назад.
		ptr := i % k

		// --- Логика Динамического Программирования ---

		// 1. Ищем тройку (min3)
		// Мы используем текущее val как *третий* элемент.
		// Нам нужна лучшая *пара*, которую мы нашли K шагов назад.
		// Это значение m2_to_use.
		m2_to_use := min2_buffer[ptr]
		if m2_to_use != infinity {
			// Нашли кандидата на тройку: (пара K шагов назад) + (текущее число)
			global_min3 = min_int64(global_min3, m2_to_use+val)
		}

		// 2. Ищем пару (min2)
		// Мы используем текущее val как *второй* элемент.
		// Нам нужен лучший *первый* элемент, который мы нашли K шагов назад.
		// Это значение m1_to_use.
		m1_to_use := min1_buffer[ptr]
		if m1_to_use != infinity {
			// Нашли кандидата на пару: (первый K шагов назад) + (текущее число)
			// Обновляем *глобальный* минимум для пары
			global_min2 = min_int64(global_min2, m1_to_use+val)
		}

		// 3. Ищем первый элемент (min1)
		// Обновляем *глобальный* минимум для одного числа
		global_min1 = min_int64(global_min1, val)

		// 4. Сохраняем текущие лучшие значения в буфер
		// В ячейку, на которую указывает ptr, мы записываем
		// лучшие *на данный момент* global_min1 и global_min2.
		// Через K итераций эти значения будут "вытащены"
		// как m1_to_use и m2_to_use.
		min1_buffer[ptr] = global_min1
		min2_buffer[ptr] = global_min2
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Ошибка при чтении файла %s: %v", filename, err)
	}

	// Если мы не нашли ни одной тройки, global_min3 останется 'infinity'

	if global_min3 == infinity {
		return -1 // или 0, в зависимости от того, как обрабатывать отсутствие ответа
	}

	return global_min3
}
