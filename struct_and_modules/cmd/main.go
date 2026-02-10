package main

import (
	"basics/struct_and_modules/pkg/handler"
	"basics/struct_and_modules/pkg/utils"
	"fmt"
)

func main() {
	fmt.Println(utils.CountChars(""))        // вернет 5
	fmt.Println(utils.CountChars("Привет"))  // вернет 12
	fmt.Println(utils.CountChars("Hello 👋")) // вернет 10

	fmt.Println(handler.ProcessOrder("Laptop", 1200.50, 1))
	fmt.Println(handler.ProcessOrder("Mouse", 25.00, 2))
}
