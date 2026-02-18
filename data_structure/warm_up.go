package main

import (
	"basics/data_structure/warm_up"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(warm_up.AppendUnique(nums, 2)) // [1 2 3]
	fmt.Println(warm_up.AppendUnique(nums, 4)) // [1 2 3 4]

	fmt.Println(warm_up.CutTail(nums, 2))  // [1]
	fmt.Println(warm_up.CutTail(nums, 5))  // []
	fmt.Println(warm_up.CutTail(nums, -1)) // []

	m := map[string]int{"a": 1, "b": 2, "c": 1}
	fmt.Println(warm_up.KeyExists(m, "a")) // true
	fmt.Println(warm_up.KeyExists(m, "z")) // false

	fmt.Println(warm_up.CountValues(m)) // map[1:2 2:1]
}
