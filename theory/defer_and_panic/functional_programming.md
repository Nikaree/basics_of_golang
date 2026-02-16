# Функциональное программирование в Go
**Функциональное программирование** — это подход, при котором основной акцент делается на использовании функций, 
как базовых строительных блоков программы. В отличие от императивного программирования, где упор делается на изменение
состояния и выполнение последовательности команд, функциональный подход сосредоточен на вычислении результатов 
через функции.

В Go поддерживается ряд элементов функционального программирования, таких как функции высшего порядка и замыкания, 
хотя язык не является чисто функциональным. Go предлагает прагматичный подход, позволяя комбинировать функциональные 
концепции с императивным стилем.

## 1. Функции как переменные
   В Go функции можно присваивать переменным, что делает код гибким и переиспользуемым:
```go
func add(a, b int) int {
	return a + b
}

func main() {
	sum := add  // присваиваем функцию переменной
	fmt.Println(sum(3, 4)) // выводит 7
	
	// Можно даже создавать срезы или мапы функций
	operations := map[string]func(int, int) int{
		"add":      add,
		"subtract": func(a, b int) int { return a - b },
		"multiply": func(a, b int) int { return a * b },
	}
	
	fmt.Println(operations["add"](5, 3))      // выводит 8
	fmt.Println(operations["subtract"](5, 3)) // выводит 2
	fmt.Println(operations["multiply"](5, 3)) // выводит 15
}
```
## 2. Функции как аргументы
   Функции могут принимать другие функции в качестве аргументов, что позволяет создавать универсальные алгоритмы:
```go
func operate(a, b int, operation func(int, int) int) int {
return operation(a, b)
}

func multiply(a, b int) int {
return a * b
}

func main() {
result := operate(3, 4, multiply)
fmt.Println(result) // выводит 12

	// Можно также использовать анонимную функцию
	sum := operate(3, 4, func(x, y int) int {
		return x + y
	})
	fmt.Println(sum) // выводит 7
}
```
## 3. Замыкания (closures)
   Это функции, которые могут использовать переменные из окружающего их контекста:
```go
func counter() func() int {
	count := 0  // переменная существует только в рамках counter
	return func() int {
		count++  // но доступна для возвращаемой функции
		return count
	}
}

func main() {
	inc := counter()
	fmt.Println(inc()) // выводит 1
	fmt.Println(inc()) // выводит 2
	
	// Создаем новый счетчик
	inc2 := counter()
	fmt.Println(inc2()) // выводит 1
	
	// Первый счетчик сохраняет свое состояние
	fmt.Println(inc()) // выводит 3
}
```
Замыкания удобны для сохранения состояния между вызовами функций.
## 4. Функции высшего порядка
   Это функции, которые либо принимают другие функции в качестве аргументов, либо возвращают функции как результат:
```go
func makeAdder(x int) func(int) int {
return func(y int) int {
return x + y
}
}

func main() {
addFive := makeAdder(5)
addTen := makeAdder(10)

	fmt.Println(addFive(3)) // выводит 8
	fmt.Println(addTen(3))  // выводит 13
}
```
## 5. Функциональные итераторы в Go
   Хотя Go не имеет встроенных функций map, filter и reduce, как в некоторых других языках, мы можем легко реализовать их сами:
```go
// Map применяет функцию f к каждому элементу среза и возвращает новый срез
func Map(data []int, f func(int) int) []int {
	result := make([]int, len(data))
	for i, v := range data {
		result[i] = f(v)
	}
	return result
}

// Filter возвращает срез, содержащий только те элементы, для которых f возвращает true
func Filter(data []int, f func(int) bool) []int {
	var result []int
	for _, v := range data {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce применяет функцию f к элементам среза, агрегируя их в один результат
func Reduce(data []int, initialValue int, f func(int, int) int) int {
	result := initialValue
	for _, v := range data {
		result = f(result, v)
	}
	return result
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	
	// Пример использования Map
	squares := Map(data, func(x int) int {
		return x * x
	})
	fmt.Println("Квадраты:", squares) // [1 4 9 16 25]
	
	// Пример использования Filter
	even := Filter(data, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Четные числа:", even) // [2 4]
	
	// Пример использования Reduce
	sum := Reduce(data, 0, func(acc, x int) int {
		return acc + x
	})
	fmt.Println("Сумма:", sum) // 15
}
```
## 6. Чистые функции
   Чистые функции — важная концепция функционального программирования:
   * Они всегда возвращают одинаковый результат для одинаковых аргументов
   * Не имеют побочных эффектов (не изменяют внешнее состояние)
```go
// Чистая функция
func add(a, b int) int {
	return a + b
}

// Нечистая функция (имеет побочный эффект)
var total = 0
func addToTotal(value int) int {
	total += value  // побочный эффект
	return total
}
```
Стремление к написанию чистых функций делает код более предсказуемым и тестируемым.




