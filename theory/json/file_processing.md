# 📌 Как работать с файлами в Go
В Go работа с файлами происходит через пакет:
```go
import "os"
```
Основные операции:
* Создать файл → os.Create
* Открыть файл → os.Open
* Открыть для чтения/записи → os.OpenFile
* Удалить → os.Remove
* Проверить существование → os.Stat

# 📌 Открытие и закрытие файлов
Файлы открываются с помощью функции os.Open() и обязательно должны быть закрыты после работы с помощью метода Close().
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// дальнейшая работа с файлом
}
```
Использование defer file.Close() гарантирует, что файл будет закрыт даже при возникновении ошибок.

# 📌 Чтение файлов
## ✅ Простой способ — прочитать весь файл целиком
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Читаем весь файл в память
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	// data — это []byte
	fmt.Println(string(data))
}
```
### 🔹 Что происходит:
* os.ReadFile открывает файл
* читает его полностью
* закрывает автоматически
* возвращает []byte

⚠ Подходит только для небольших файлов.

## ✅ Чтение построчно (эффективнее)
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Ошибка открытия:", err)
		return
	}
	defer file.Close() // гарантированное закрытие

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
	}
}
```
### 🔹 Зачем так?
* Не грузим файл целиком в память
* Читаем по строкам
* Подходит для больших логов

# 📌 Запись в файл
Флаги:
* `os.WriteFile` — простая запись 
* `os.O_APPEND` — писать в конец
* `os.O_WRONLY` — только запись
* `os.O_RDWR` — чтение и запись
* `os.O_CREATE` — создать, если нет
```go
package main

import (
	"fmt"
	"os"
)

func main() {

	// Открываем (или создаём) файл example.txt
	// os.O_CREATE  — создать файл, если его нет
	// os.O_RDWR    — открыть для чтения и записи
	// os.O_APPEND  — все записи будут добавляться в конец файла
	file, err := os.OpenFile(
		"example.txt",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0644, // права доступа: владелец может читать и писать, остальные — только читать
	)
	if err != nil {
		panic(err)
	}
	defer file.Close() // обязательно закрываем файл

	// 🔹 Записываем строку (она добавится в конец файла)
	_, err = file.WriteString("Привет, это новая строка\n")
	if err != nil {
		panic(err)
	}

	// 🔹 Перед чтением перемещаем курсор в начало файла
	// (иначе будем читать с текущей позиции — с конца)
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	// 🔹 Читаем содержимое файла
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}

	// Выводим прочитанные данные
	fmt.Println("Содержимое файла:")
	fmt.Println(string(buffer[:n]))
}

```

## ✅ Создание и запись
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка создания:", err)
		return
	}
	defer file.Close()

	// Записываем строку
	_, err = file.WriteString("Привет, Go!\n")
	if err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}

	fmt.Println("Запись успешна")
}
```

## ✅ Добавление в конец файла (append)
```go
file, err := os.OpenFile(
	"output.txt",
	os.O_APPEND|os.O_WRONLY,
	0644,
)
```

# 📌 Итоговая таблица
| Что                     | Как                       |
| ----------------------- | ------------------------- |
| Прочитать весь файл     | `os.ReadFile`             |
| Читать по строкам       | `bufio.Scanner`           |
| Создать файл            | `os.Create`               |
| Добавить в конец        | `os.OpenFile + O_APPEND`  |
| Закрыть файл            | `defer file.Close()`      |
| Проверить существование | `os.Stat + os.IsNotExist` |

# 📌 Управление ресурсами (ОЧЕНЬ важно)
Когда ты открываешь файл:
```go
file, err := os.Open("file.txt")
```
Ты получаешь системный ресурс (file descriptor).

**Если не закрыть файл:**
* будет утечка ресурсов
* программа может перестать открывать новые файлы
* возможны проблемы в продакшене

✅ Правильный способ
```go
defer file.Close()
```
**defer** гарантирует закрытие даже при panic или return.

# 📌 Обработка ошибок при работе с файлами
В Go ошибки — это обычные значения.

## ✅ Проверка существования файла
```go
_, err := os.Stat("file.txt")

if os.IsNotExist(err) {
    fmt.Println("Файл не существует")
}
```

## ✅ Различение ошибок
```go
file, err := os.Open("file.txt")
if err != nil {
    if os.IsNotExist(err) {
        fmt.Println("Файл не найден")
    } else {
        fmt.Println("Другая ошибка:", err)
    }
return
}
```

# 📌 Работа с директориями 
## ✅ os.Mkdir -> Создание директории
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Создаём папку с правами 0755
	err := os.Mkdir("mydir", 0755)
	if err != nil {
		fmt.Println("Ошибка создания:", err)
		return
	}

	fmt.Println("Директория создана")
}
```
### 🔹 Что означает 0755?
Это права доступа:
```css
0 - восьмеричная система
7 - владелец (rwx)
5 - группа (r-x)
5 - остальные (r-x)
```

## ✅ os.MkdirAll -> Создание вложенных директорий
```go
// Если нужно создать: a/b/c
// И папки a и b ещё не существуют — обычный Mkdir не сработает.
err := os.MkdirAll("a/b/c", 0755)
if err != nil {
	fmt.Println("Ошибка:", err)
	return
}
```
### 🔹 MkdirAll:
* создаёт все недостающие папки
* не выдаёт ошибку, если директория уже существует

## ✅ os.Remove -> Удаление пустой директории
```go
err := os.Remove("mydir")
if err != nil {
	fmt.Println("Ошибка удаления:", err)
}
```

**⚠ Важно:**
* Удаляет только пустую директорию
* Если внутри есть файлы — будет ошибка

## ✅ os.RemoveAll -> Удаление директории с содержимым
```go
err := os.RemoveAll("a")
if err != nil {
	fmt.Println("Ошибка удаления:", err)
}
```
🔹 **Удаляет:**
* папку
* все вложенные файлы
* все поддиректории
**⚠ Очень опасная функция — может удалить всё дерево.**

# 📌 Проверка существования директории
```go
info, err := os.Stat("mydir")

if os.IsNotExist(err) {
	fmt.Println("Директория не существует")
	return
}

if info.IsDir() {
	fmt.Println("Это директория")
}
```