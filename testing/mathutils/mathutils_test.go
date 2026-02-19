package mathutils

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	// TODO: Определите тестовые случаи в виде таблицы (слайс структур)
	testCases := []struct {
		name     string // Имя под-теста
		a        int    // Входное значение 1
		b        int    // Входное значение 2
		expected int    // Ожидаемый результат
	}{
		// Добавьте несколько тестовых случаев:
		// - Сложение положительных чисел
		// - Сложение отрицательных чисел
		// - Сложение с нулем
		// - Сложение положительного и отрицательного
		{"Сумма положительных:", 2, 3, 5},
		{"Сумма отрицательных:", -2, -3, -5},
		{"Сумма разнознаковых:", 2, -3, -1},
		{"Сумма с нулем:", 2, 0, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// TODO: Вызовите тестируемую функцию Add
			result := Add(tc.a, tc.b)

			// TODO: Сравните результат с ожидаемым (tc.expected)
			// Используйте t.Errorf() для сообщения об ошибке, если результат не совпадает.
			// Пример: if result != tc.expected { t.Errorf("ожидали %d, получили %d", tc.expected, result) }
			if result != tc.expected {
				t.Errorf("ожидали %d, получили %d", tc.expected, result)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	// TODO: Напишите тесты для функции Subtract, используя табличный подход.
	// Не забудьте проверить разные комбинации чисел (положительные, отрицательные, ноль).
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Вычитание положительных", 5, 3, 2},
		{"Вычитание отрицательных", -5, -3, -2},
		{"Вычитание разнознаковых", -5, 3, -8},
		{"Вычитание положительного с нулем", 5, 0, 5},
		{"Вычитание отрицательного с нулем", -5, 0, -5},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Subtract(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("ожидали %d, получили %d", tc.expected, result)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	// TODO: Напишите тесты для функции Factorial, используя табличный подход.
	testCases := []struct {
		name             string
		n                int
		expectedVal      int
		expectedErr      bool   // true, если ожидается ошибка, иначе false
		errorMsgContains string // Опционально: подстрока для проверки в тексте ошибки
	}{
		// Проверьте случаи:
		// - n = 0 (ожидаем 1, ошибки нет)
		// - n = 1 (ожидаем 1, ошибки нет)
		// - n = 5 (ожидаем 120, ошибки нет)
		// - n = -1 (ожидаем ошибку, значение не важно)
		{"Факториал нуля:", 0, 1, false, "not"},
		{"Факториал пяти:", 5, 120, false, "not"},
		{"Факториал отрицательного:", -1, 0, true, "must be positive"}, // Ожидаем ошибку
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := Factorial(tc.n)

			if tc.expectedErr {
				if err == nil {
					t.Fatalf("ожидали ошибку, но получили nil")
				}
				if !strings.Contains(err.Error(), tc.errorMsgContains) {
					t.Errorf("ожидали: %v, но получили: %v", tc.errorMsgContains, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("не ожидали ошибку, но получили: %v", err)
				}

				if result != tc.expectedVal {
					t.Errorf("ожидали %d, получили %d", tc.expectedVal, result)
				}
			}
		})
	}
}
