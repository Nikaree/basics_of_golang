//Создайте пакет temperature с файлом temperature.go.
//Реализуйте тип Temperature, хранящий значение температуры во внутреннем формате.
//Предоставьте три конструктора: NewCelsius, NewFahrenheit и NewKelvin.
//Реализуйте методы Celsius(), Fahrenheit() и Kelvin() для конвертации.
//Реализуйте метод String() для соответствия интерфейсу fmt.Stringer.

package main

import (
	"basics/type_for_temp/temperature"
	"fmt"
)

func main() {
	t1 := temperature.NewCelsius(25.0)
	fmt.Println(t1)                                      // "25.0°C"
	fmt.Printf("В Фаренгейтах: %.1f\n", t1.Fahrenheit()) // "В Фаренгейтах: 77.0"

	t2 := temperature.NewFahrenheit(98.6)
	fmt.Println(t2)                                      // "37.0°C"
	fmt.Printf("В Фаренгейтах: %.1f\n", t2.Fahrenheit()) // "В Фаренгейтах: 77.0"

	t3 := temperature.NewKelvin(0)
	fmt.Println(t3) // "-273.1°C"
}
