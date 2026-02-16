package main

import (
	"basics/defer_and_panic/fp"
	"fmt"
)

func main() {
	nums := []int{-2, -1, 0, 1, 2, 3}

	// Использование вспомогательных функций
	positive := fp.FilterPositive(nums)
	fmt.Println("Положительные числа:", positive) // [1 2 3]

	squared := fp.SquareAll(nums)
	fmt.Println("Квадраты чисел:", squared) // [4 1 0 1 4 9]

	sum := fp.Sum(nums)
	fmt.Println("Сумма всех чисел:", sum) // 3

	// Использование HOF функций напрямую
	even := fp.FilterInt(nums, func(x int) bool { return x%2 == 0 })
	fmt.Println("Чётные числа:", even) // [-2 0 2]
}
