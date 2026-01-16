package EGEINFO

import (
	"bufio"   // Для эффективного чтения файла построчно
	"fmt"     // Для вывода в консоль
	"log"     // Для вывода фатальных ошибок
	"os"      // Для работы с файлами
	"sort"    // Для сортировки
	"strconv" // Для конвертации строк в числа (Atoi)
	"strings" // Для разделения строки (Split)
)

func Scan_FileEge1_26(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Ошибка: Не могу открыть файл '%s': %v", filename, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		log.Fatal("Ошибка: Файл пуст или не удалось прочитать первую строку")
	}

	// Разделяем строку по пробелу (аналог .split() в Python)
	firstLineParts := strings.Split(scanner.Text(), " ")
	if len(firstLineParts) != 2 {
		log.Fatal("Ошибка: Первая строка должна содержать два числа (S и N)")
	}

	// Конвертируем строки в числа (аналог int())
	S, err := strconv.Atoi(firstLineParts[0])
	if err != nil {
		log.Fatalf("Ошибка: Не удалось S ('%s') в число: %v", firstLineParts[0], err)
	}

	N, err := strconv.Atoi(firstLineParts[1])
	if err != nil {
		log.Fatalf("Ошибка: Не удалось N ('%s') в число: %v", firstLineParts[1], err)
	}

	// 3. Читаем N размеров файлов
	// Создаем 'слайс' (динамический массив) для хранения размеров.
	// Мы сразу указываем N в 'cap' (capacity) для оптимизации,
	// чтобы Go не выделял память заново при каждом добавлении.
	files := make([]int, 0, N)

	for i := 0; i < N; i++ {
		if !scanner.Scan() {
			// Ошибка, если файл закончился раньше, чем мы ожидали
			log.Fatalf("Ошибка: Ожидалось %d файлов, но файл закончился после %d", N, i)
		}

		fileSize, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Ошибка: Не удалось прочитать размер файла на строке %d: %v", i+2, err)
		}
		// Добавляем размер в наш слайс (аналог .append() в Python)
		files = append(files, fileSize)
	}

	// Обязательно проверяем, не было ли ошибок во время сканирования
	if err := scanner.Err(); err != nil {
		log.Fatal("Ошибка при чтении файла:", err)
	}

	// --- Начало алгоритма ---

	// 1. Сортируем все файлы по возрастанию (аналог files.sort())
	sort.Ints(files)

	// 2. Жадный алгоритм
	savedFilesSum := 0
	savedFilesCount := 0
	lastSavedFileIndex := -1 // Индекс последнего "влезшего" файла

	for i := 0; i < N; i++ {
		if savedFilesSum+files[i] <= S {
			savedFilesSum += files[i]
			savedFilesCount++
			lastSavedFileIndex = i
		} else {
			// Этот файл и все следующие уже не влезут
			break
		}
	}

	// Обработка случая, если ни один файл не влез
	if savedFilesCount == 0 {
		fmt.Println(0)
		fmt.Println(0)
		return // Завершаем программу
	}

	// Первый ответ — максимальное число пользователей
	fmt.Println(savedFilesCount)

	// 3. Ищем второй ответ

	// Сумма всех, КРОМЕ самого большого из сохраненных
	sumWithoutLast := savedFilesSum - files[lastSavedFileIndex]

	// По умолчанию, самый большой файл, который мы можем взять
	maxPossibleFile := files[lastSavedFileIndex]

	// Идём по файлам, которые мы НЕ взяли
	for i := lastSavedFileIndex + 1; i < N; i++ {
		fileToCheck := files[i]

		if sumWithoutLast+fileToCheck <= S {
			// Да, влезли. Обновляем кандидата
			maxPossibleFile = fileToCheck
		} else {
			// Этот файл 'fileToCheck' уже не влез.
			// Все следующие тем более не влезут.
			break
		}
	}

	// Второй ответ - максимальный размер файла
	fmt.Println(maxPossibleFile)
}
