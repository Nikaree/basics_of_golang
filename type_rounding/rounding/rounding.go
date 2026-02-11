package rounding

import (
	"fmt"
	"math"
	"strconv"
)

// FloorRounder реализует округление вниз (к меньшему числу).
type FloorRounder struct{}

// CeilRounder реализует округление вверх (к большему числу).
type CeilRounder struct{}

// MathRounder реализует математическое округление.
type MathRounder struct{}

// Round округляет число вниз до указанного количества знаков после запятой.
func (f FloorRounder) Round(value float64, places int) float64 {
	if places < 0 {
		places = 0
	}

	factor := math.Pow(10, float64(places))
	return math.Floor(value*factor) / factor
}

// Round округляет число вверх до указанного количества знаков после запятой.
func (c CeilRounder) Round(value float64, places int) float64 {
	if places < 0 {
		places = 0
	}

	factor := math.Pow(10, float64(places))
	return math.Ceil(value*factor) / factor
}

// Round выполняет математическое округление (0.5 округляется от нуля).
func (m MathRounder) Round(value float64, places int) float64 {
	if places < 0 {
		places = 0
	}

	factor := math.Pow(10, float64(places))
	return math.Round(value*factor) / factor
}

// FormatFloat форматирует число с точно указанным количеством знаков после запятой.
func FormatFloat(value float64, places int) string {
	if places < 0 {
		places = 0
	}

	return fmt.Sprintf("%.*f", places, value)
}

// StringToInt преобразует строку в int.
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// TruncateToInt отбрасывает дробную часть (округление к нулю).
func TruncateToInt(value float64) int {
	return int(value)
}
