package main

import (
	"basics/memory_ptr/typed_payments/payments"
	"fmt"
)

func main() {
	// Создаем каталог товаров
	products := map[payments.ProductID]payments.Product{
		"apple": {
			ID:        "apple",
			Name:      "Яблоко",
			Price:     payments.Money(50), // 50 центов = $0.50
			Available: payments.Count(100),
		},
		"banana": {
			ID:        "banana",
			Name:      "Банан",
			Price:     payments.Money(30), // 30 центов = $0.30
			Available: payments.Count(50),
		},
	}

	// Создаем корзину
	cart := payments.Cart{Items: make(map[payments.ProductID]payments.Count)}

	// Добавляем товары в корзину
	fmt.Println("Добавляем 5 яблок...")
	err := cart.AddProduct("apple", payments.Count(5))
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println("Добавляем 3 банана...")
	err = cart.AddProduct("banana", payments.Count(3))
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println("Пытаемся добавить 0 яблок...")
	err = cart.AddProduct("apple", payments.Count(0))
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err) // Должна быть ошибка
	}

	// Вычисляем общую стоимость
	total := cart.TotalPrice(products)
	// Ожидаем: (5 * 50) + (3 * 30) = 250 + 90 = 340
	fmt.Println("Общая стоимость:", total)                   // Должно вывести $3.40
	fmt.Println("Общая стоимость (формат):", total.String()) // Вывод через метод String()
}
