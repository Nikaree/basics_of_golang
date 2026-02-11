package main

import (
	"basics/pointer_mutation/pointers"
	"fmt"
)

func main() {
	// Тест Mutate
	num := 5
	err := pointers.Mutate(&num)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(num) // Должно вывести 15
	}

	// Тест ReverseString
	str := "Hello, мир!"
	err = pointers.ReverseString(&str)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(str) // Должно вывести "!рим ,olleH"
	}
}
