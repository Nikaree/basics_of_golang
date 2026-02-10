package handler

import (
	"basics/struct_and_modules/internal/calculator"
	"fmt"
)

func ProcessOrder(item string, price float64, qty int) (string, string) {

	total := calculator.CalculateTotal(price, qty)
	line1 := fmt.Sprintf("Processing order for %d x %s", qty, item)
	line2 := fmt.Sprintf("$%.2f with 2 decimals", total)

	return line1, line2
}
