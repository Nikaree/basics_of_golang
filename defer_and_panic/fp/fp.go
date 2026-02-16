package fp

// FilterInt возвращает новый срез, содержащий только те элементы slice,
// для которых pred(x) == true
func FilterInt(slice []int, pred func(int) bool) []int {
	var out []int
	for _, v := range slice {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}

// MapInt применяет mapper(x) к каждому элементу и возвращает новый срез с результатами
func MapInt(slice []int, mapper func(int) int) []int {
	var out []int
	for _, v := range slice {
		out = append(out, mapper(v))
	}
	return out
}

// ReduceInt сворачивает срез в одно значение, используя функцию reducer
// initial — начальное значение аккумулятора
func ReduceInt(slice []int, initial int, reducer func(acc, x int) int) int {
	acc := initial
	for _, v := range slice {
		acc = reducer(acc, v)
	}
	return acc
}

// FilterPositive оставляет только положительные числа (> 0)
// ДОЛЖНА использовать FilterInt
func FilterPositive(nums []int) []int {
	out := FilterInt(nums, func(x int) bool {
		return x > 0
	})
	return out
}

// SquareAll возводит каждое число в квадрат
// ДОЛЖНА использовать MapInt
func SquareAll(nums []int) []int {
	out := MapInt(nums, func(x int) int {
		return x * x
	})
	return out
}

// Sum считает сумму всех чисел
// ДОЛЖНА использовать ReduceInt
func Sum(nums []int) int {
	sum := ReduceInt(nums, 0, func(acc, x int) int {
		return acc + x
	})
	return sum
}
