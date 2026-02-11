package temperature

import (
	"fmt"
	"math"
)

const absoluteZeroC = -273.15
const absoluteZeroF = -459.67
const absoluteZeroK = 0

type Temperature struct {
	temp float64
}

func clamToAbsoluteZero(t float64, absoluteZero float64) float64 {
	if t < absoluteZero {
		return absoluteZero
	}
	return t
}

// NewCelsius создаёт температуру из градусов Цельсия
func NewCelsius(t float64) Temperature {
	return Temperature{temp: clamToAbsoluteZero(t, absoluteZeroC)}
}

// NewFahrenheit создаёт температуру из градусов Фаренгейта
func NewFahrenheit(t float64) Temperature {
	с := (t - 32) * 5 / 9
	return Temperature{temp: clamToAbsoluteZero(с, absoluteZeroF)}
}

// NewKelvin создаёт температуру из Кельвинов
func NewKelvin(t float64) Temperature {
	return Temperature{temp: clamToAbsoluteZero(t, absoluteZeroK) - 273.15}
}

// Celsius Конвертация в градусы Цельсия
func (t Temperature) Celsius() float64 {
	return t.temp
}

// Fahrenheit Конвертация в градусы Фаренгейта
func (t Temperature) Fahrenheit() float64 {
	return t.temp*9/5 + 32
}

// Kelvin Конвертация в градусы Кельвина
func (t Temperature) Kelvin() float64 {
	return t.temp - 273.15
}

// String возвращает строку вида "25.0°C" (везде где ты пишешь используя fmt.Stringer)
func (t Temperature) String() string {
	rounded := math.Round(t.temp*10) / 10
	return fmt.Sprintf("%.1f°C", rounded)
}
