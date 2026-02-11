package payments

import (
	"fmt"
)

type Money int        // Для представления денежных сумм в центах
type Count int        // Для представления количества товаров
type ProductID string // Для представления идентификаторов товаров

var (
	ErrCount = fmt.Errorf("count must be greater than zero")
)

// String Форматирует Money (центы) в строку вида "$X.XX" для положительных сумм и нуля,
// и "$-X.XX" для отрицательных (например, Money(150) → "$1.50", Money(-75) → "$-0.75"). Используйте fmt.Sprintf.
func (m Money) String() string {
	sign := ""
	if m < 0 {
		sign = "-"
	}

	dollars := m / 100
	cents := m % 100

	return fmt.Sprintf("$%s%d.%02d", sign, dollars, cents)
}

// Add Складывает две денежные суммы.
func (m Money) Add(other Money) Money {
	return m + other
}

// Subtract Вычитает одну денежную сумму из другой.
func (m Money) Subtract(other Money) Money {
	return m - other
}

// Multiply Умножает денежную сумму (например, цену) на количество товаров.
func (m Money) Multiply(count Count) Money {
	return m * Money(count)
}

type Product struct {
	ID        ProductID
	Name      string
	Price     Money // Цена в центах
	Available Count // Доступное количество на складе
}
type Cart struct {
	// Ключ - ProductID, Значение - количество (Count) товара в корзине
	Items map[ProductID]Count
}

// AddProduct Добавляет товар с указанным id в корзину в количестве count.
// Если count меньше или равен нулю, метод должен возвращать ошибку (можно использовать errors.New или fmt.Errorf).
// Если товара с таким id уже есть в корзине, количество должно увеличиться.
// Если корзина (c.Items) еще не инициализирована (равна nil), ее нужно инициализировать с помощью make.
// Примечание: Этот метод, согласно ТЗ, не проверяет наличие товара в каталоге или его доступное количество на складе.
func (c *Cart) AddProduct(id ProductID, count Count) error {
	if count <= 0 {
		return fmt.Errorf("error with add product: %w", ErrCount)
	}
	if c.Items == nil {
		c.Items = make(map[ProductID]Count)
	}
	c.Items[id] += count
	return nil
}

// TotalPrice Вычисляет и возвращает общую стоимость всех товаров в корзине.
// Для расчета используются цены из переданной карты products (каталога товаров).
// Если какой-то товар из корзины (c.Items) отсутствует в каталоге products, он должен быть
// проигнорирован при подсчете общей стоимости.
func (c *Cart) TotalPrice(products map[ProductID]Product) Money {
	var total Money

	if c.Items == nil {
		return total
	}

	for id, count := range c.Items {
		product, ok := products[id]
		if ok {
			total += product.Price.Multiply(count)
		}
	}
	return total
}
