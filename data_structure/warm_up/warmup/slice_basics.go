package warmup

// AppendUnique Возвращает новый срез с v, если его нет в nums; исходный nums не менять.
// Важно: новый срез должен иметь собственный массив (отдельный backing-array).
func AppendUnique(nums []int, v int) []int {
	for _, n := range nums {
		if n == v {
			return nums
		}
	}
	return append(nums, v)
}

// CutTail Возвращает копию первых len(nums) - n элементов.
// Если n >= len(nums) или n < 0 — вернуть пустой срез ([]int{}).
func CutTail(nums []int, n int) []int {
	if n >= len(nums) || n < 0 {
		return []int{}
	}
	result := make([]int, 0, n)

	index := len(nums) - n
	for _, v := range nums[:index] {
		result = append(result, v)
	}
	return result
}
