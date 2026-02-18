package warm_up

// KeyExists Возвращает true, если ключ есть в карте, иначе false.
func KeyExists(m map[string]int, key string) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

// CountValues Возвращает карту, где ключ — значение из m, а значение — количество раз, сколько оно встречается.
func CountValues(m map[string]int) map[int]int {
	countVals := make(map[int]int)
	for _, v := range m {
		countVals[v]++
	}
	return countVals
}
