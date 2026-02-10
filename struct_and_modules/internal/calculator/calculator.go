package calculator

func CalculateTotal(price float64, quantity int) float64 {
	result := price * float64(quantity) * 1.10
	return result
}
