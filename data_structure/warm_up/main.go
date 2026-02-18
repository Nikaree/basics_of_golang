package main

import (
	"basics/data_structure/warm_up/warmup"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(warmup.AppendUnique(nums, 2)) // [1 2 3]
	fmt.Println(warmup.AppendUnique(nums, 4)) // [1 2 3 4]

	fmt.Println(warmup.CutTail(nums, 2))  // [1]
	fmt.Println(warmup.CutTail(nums, 5))  // []
	fmt.Println(warmup.CutTail(nums, -1)) // []

	m := map[string]int{"a": 1, "b": 2, "c": 1}
	fmt.Println(warmup.KeyExists(m, "a")) // true
	fmt.Println(warmup.KeyExists(m, "z")) // false

	fmt.Println(warmup.CountValues(m)) // map[1:2 2:1]
}
