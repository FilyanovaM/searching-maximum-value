package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements создаёт слайс случайных положительных чисел.
// При неположительном размере возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)

	for i := range data {
		data[i] = rand.Intn(1_000_000) + 1
	}

	return data
}

// maximum возвращает максимальный элемент слайса.
// Для пустого слайса и nil возвращает 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]

	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks находит максимум, разделяя слайс на CHUNKS частей.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	chunkMaxima := make([]int, CHUNKS)

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		// Последняя часть включает остаток от деления.
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)

		go func(index int, chunk []int) {
			defer wg.Done()

			// При len(data) < CHUNKS некоторые части пустые.
			// Используем элемент исходного слайса, чтобы пустая
			// часть не повлияла на результат даже для отрицательных чисел.
			if len(chunk) == 0 {
				chunkMaxima[index] = data[0]
				return
			}

			chunkMaxima[index] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()

	return maximum(chunkMaxima)
}

func main() {
	fmt.Printf("Генерируем %d целых положительных чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение последовательно")

	start := time.Now()
	sequentialMax := maximum(data)
	sequentialElapsed := time.Since(start).Microseconds()

	fmt.Printf(
		"Максимальное значение элемента: %d\nВремя поиска: %d мкс\n",
		sequentialMax,
		sequentialElapsed,
	)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)

	start = time.Now()
	parallelMax := maxChunks(data)
	parallelElapsed := time.Since(start).Microseconds()

	fmt.Printf(
		"Максимальное значение элемента: %d\nВремя поиска: %d мкс\n",
		parallelMax,
		parallelElapsed,
	)

	if sequentialMax != parallelMax {
		fmt.Println("Ошибка: результаты поиска не совпадают")
	}
}
