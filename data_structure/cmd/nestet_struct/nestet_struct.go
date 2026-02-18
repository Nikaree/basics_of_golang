// Создайте вложенную структуру, описывающую данные книги (автор, издательство, год выпуска).
// Реализуйте метод для форматированного вывода данных книги.
// Напишите пример сериализации и десериализации структуры книги в JSON.

package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name string `json:"author_name"`
}
type Publisher struct {
	Name string `json:"publisher_name"`
	Year int    `json:"publisher_year"`
}
type Book struct {
	Title     string `json:"title"`
	Author    `json:"author"`
	Publisher `json:"publisher"`
}

func (b Book) Print() string {
	return fmt.Sprintf("'%s' by %s, published by %s in %d",
		b.Title, b.Author.Name, b.Publisher.Name, b.Publisher.Year)
}

func main() {
	book := Book{"Go Programming", Author{"Alan Turing"}, Publisher{"Penguin", 2020}}
	fmt.Println(book.Print())

	jsonData, _ := json.Marshal(book)
	fmt.Println(string(jsonData))

	var decodedBook Book
	json.Unmarshal(jsonData, &decodedBook)
	fmt.Println(decodedBook.Print())
}
