**JSON (JavaScript Object Notation)** — это текстовый формат хранения и передачи данных.

Пример JSON:
```json
{
"name": "Anna",
"age": 26,
"is_student": false
}
```
**🔹 Почему его используют?**
* Лёгкий для чтения человеком
* Удобный для передачи по HTTP
* Поддерживается почти всеми языками

# 📌 Работа с JSON в Go
В Go используется пакет:
```go
import "encoding/json"
```
Есть две основные операции:
* Marshal → Go → JSON
* Unmarshal → JSON → Go

# 📌 Кодирование (Marshal)
## ✅ Go → JSON
```go
package main

import (
	"encoding/json" // пакет для работы с JSON
	"fmt"
)

// Структура должна иметь экспортируемые поля (с большой буквы)
type User struct {
	Name string `json:"name"` // json-тег задаёт имя поля в JSON
	Age  int    `json:"age"`
}

func main() {
	// Создаём экземпляр структуры
	user := User{
		Name: "Anna",
		Age:  26,
	}

	// Marshal преобразует структуру Go в JSON (в []byte)
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}

	// Преобразуем []byte в строку для вывода
	fmt.Println(string(data))
}
```
## Результат
```json
{"Name":"Anna","Age":26}
```

# 📌 Декодирование (Unmarshal)
## ✅ JSON → Go
```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// JSON в виде строки
	jsonData := `{"name":"Anna","age":26}`

	var user User

	// Unmarshal преобразует JSON ([]byte) в структуру
	// ВАЖНО: передаём указатель &user
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	// Теперь структура заполнена данными из JSON
	fmt.Println("Имя:", user.Name)
	fmt.Println("Возраст:", user.Age)
}
```

# 📌 JSON теги (очень важно)
Обычно API используют lowercase.
Для этого применяют теги:
```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```
Теперь результат:
```json
{"name":"Anna","age":26}
```
## 🔹 Почему это важно?
* Без тега JSON будет использовать имя поля
* Поле должно быть экспортируемым (с большой буквы!)

❌ Это не сработает:
```go
type User struct {
    name string
}
```
Потому что поле не экспортируется.

# 📌 Работа с вложенными структурами
## Файл Go
```go
type Address struct {
	City string `json:"city"`
}

type User struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}
```

## Файл json
```json
{
  "name": "Anna",
  "address": {
    "city": "Seoul"
  }
}

```

# 📌 Работа с массивами
## Файл Go
```go
type User struct {
	Name string `json:"name"`
}

users := []User{
	{Name: "Anna"},
	{Name: "Maria"},
}

data, _ := json.Marshal(users)

```
## Результат
```json
[{"name":"Anna"},{"name":"Maria"}]
```

# 📌 Практический пример — чтение JSON из файла
```go
fileData, err := os.ReadFile("user.json")
if err != nil {
	return
}

var user User
err = json.Unmarshal(fileData, &user)
if err != nil {
	return
}
```

# 📌 Оптимизация работы с JSON
## Типичные проблемы производительности:
* Аллокации памяти — json.Marshal и json.Unmarshal создают временные копии данных в памяти.
* Большие объемы данных — работа с большими JSON-файлами может вызвать проблемы с памятью.
* Повторное использование структур — каждый раз создавать новые структуры неэффективно.

## Способы оптимизации:
* Используйте потоковую обработку (json.Decoder и json.Encoder) для больших файлов.
* Пулы объектов — переиспользуйте структуры вместо их повторного создания.
* Выборочное декодирование — декодируйте только нужные поля

## 🔥 1. Использовать Encoder/Decoder вместо Marshal/Unmarshal
Это важно для больших данных:
```go
file, _ := os.Open("users.json")
defer file.Close()

decoder := json.NewDecoder(file)

var user User
decoder.Decode(&user)
```
Преимущества:
* Не загружает всё в память
* Работает потоково

## 🔥 2. Использовать omitempty
```go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}
```
Если Age == 0, поле не попадёт в JSON.