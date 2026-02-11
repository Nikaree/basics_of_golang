package escape_analysis

type LargeObject struct {
	ID   int
	Name string
	Data []int
}

// ProcessByValue Возвращает сумму всех элементов в obj.Data. Не изменяет исходный объект.
func ProcessByValue(obj LargeObject) int {
	result := 0
	for _, val := range obj.Data {
		result += val
	}
	return result
}

// ProcessByPointer Возвращает сумму всех элементов в obj.Data и увеличивает obj.ID на 1.
func ProcessByPointer(obj *LargeObject) int {
	result := 0
	for _, val := range obj.Data {
		obj.ID++
		result += val
	}
	return result
}

// CreateObjectOnStack Возвращает объект с ID=1, Name="StackObject", Data длиной 10 элементов (инициализированный нулями).
func CreateObjectOnStack() LargeObject {
	var obj LargeObject = LargeObject{
		ID:   1,
		Name: "StackObject",
		Data: make([]int, 10),
	}
	return obj
}

// CreateObjectOnHeap Возвращает указатель на объект с ID=2, Name="HeapObject",
// Data длиной 10000 элементов (инициализированный нулями).
func CreateObjectOnHeap() *LargeObject {
	var obj LargeObject = LargeObject{
		ID:   2,
		Name: "HeapObject",
		Data: make([]int, 10),
	}
	return &obj
}

func AnalyzeEscape() string {
	return "Запустите для анализа:\ngo build -gcflags=\"-m\" .\nИщите строки вида \"escapes to heap\"."
}
