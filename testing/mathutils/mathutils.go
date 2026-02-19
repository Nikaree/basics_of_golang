package mathutils

import (
	"errors"
)

// Add возвращает сумму двух целых чисел.
// TODO: Реализуйте эту функцию.
func Add(a, b int) int {
	return a + b
}

// Subtract возвращает разность двух целых чисел (a - b).
// TODO: Реализуйте эту функцию.
func Subtract(a, b int) int {
	return a - b
}

// Factorial вычисляет факториал неотрицательного целого числа n.
// Для отрицательных чисел должна возвращаться ошибка.
// TODO: Реализуйте эту функцию. Учтите граничные случаи (0! = 1).
// Используйте errors.New или fmt.Errorf для создания ошибки.
func Factorial(n int) (int, error) {
	// Ваша реализация здесь ...
	if n < 0 {
		return 0, errors.New("must be positive") // Пример ошибки
	}
	var factorial int = 1

	if n == 0 {
		return factorial, nil
	}
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial, nil // Заглушка
}
