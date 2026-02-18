package main

import "fmt"

// Deduplicate — удаление дубликатов из отсортированного среза за O(n)
func Deduplicate(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if i != len(nums)-1 && nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return nums

	/*
		if len(nums) == 0 {
			return nums
		}

		j := 1 // индекс для записи

		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] {
				nums[j] = nums[i]
				j++
			}
		}

		return nums[:j]
	*/
}

// Rotate — циклический сдвиг элементов на k позиций
func Rotate(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	k = k % n

	tmp := append(nums[n-k:], nums[:n-k]...)
	copy(nums, tmp)
	return nums
}

// Chunk — разбиение среза на подсрезы заданного размера
func Chunk(nums []int, chunkSize int) [][]int {
	rows := len(nums)/chunkSize + 1
	matrix := make([][]int, 0, rows)
	for i := 0; i < len(nums); i += chunkSize {
		if i+chunkSize > len(nums) {
			matrix = append(matrix, nums[i:])
			continue
		}
		matrix = append(matrix, nums[i:i+chunkSize])
	}
	return matrix
}

// Flatten — преобразование [][]T в []T
func Flatten(nums [][]int) []int {
	result := make([]int, 0, len(nums)*cap(nums))
	for i := range nums {
		for j := range nums[i] {
			result = append(result, nums[i][j])
		}
	}
	return result
}
func main() {
	//a := []int{1, 2, 3, 4, 5, 6}
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	//fmt.Println(Deduplicate(a))
	//fmt.Println(Rotate(a, 2))
	//fmt.Println(Chunk(a, 2))
	fmt.Println(Flatten(matrix))
}
