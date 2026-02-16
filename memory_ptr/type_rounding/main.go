package main

import (
	"basics/memory_ptr/type_rounding/rounding"
	"fmt"
)

func main() {

	// Создаем округлители
	floor := rounding.FloorRounder{}
	ceil := rounding.CeilRounder{}
	math := rounding.MathRounder{} // Используем math вместо mathRounder для краткости

	// Проверяем разные округления
	value := 3.14159

	fmt.Println("Исходное значение:", value)
	fmt.Println("Округление вниз до 2 знаков:", floor.Round(value, 2))             // 3.14
	fmt.Println("Округление вверх до 2 знаков:", ceil.Round(value, 2))             // 3.15
	fmt.Println("Математическое округление до 2 знаков:", math.Round(value, 2))    // 3.14
	fmt.Println("Математическое округление 3.15 до 1 знака:", math.Round(3.15, 1)) // 3.2

	// Проверяем форматирование
	fmt.Println("Форматирование до 3 знаков:", rounding.FormatFloat(value, 3)) // "3.142"

	// Проверяем преобразование в int
	fmt.Println("Отбрасывание дробной части:", rounding.TruncateToInt(value)) // 3

	// Преобразование строки в int
	intVal, err := rounding.StringToInt("42")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Строка в int:", intVal) // 42
	}

	_, errInvalid := rounding.StringToInt("abc")
	if errInvalid != nil {
		fmt.Println("Ожидаемая ошибка:", errInvalid)
	}

}
