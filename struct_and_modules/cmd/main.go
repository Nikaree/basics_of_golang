package main

import (
	"basics/struct_and_modules/pkg/utils"
	"fmt"
)

func main() {
	fmt.Println(utils.CountChars("Hello"))   // вернет 5
	fmt.Println(utils.CountChars("Привет"))  // вернет 12
	fmt.Println(utils.CountChars("Hello 👋")) // вернет 10
}
