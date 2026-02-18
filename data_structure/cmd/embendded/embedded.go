//Создайте тип Animal с полями Name и Sound и методом MakeSound().
//Создайте типы Dog и Cat, встраивающие тип Animal.
//Переопределите метод MakeSound() для каждого типа.
//Создайте функцию, принимающую интерфейс с методом MakeSound(), и передайте ей экземпляры Dog и Cat.

package main

import "fmt"

type Animal struct {
	Name  string
	Sound string
}
type Dog struct {
	Animal
}
type Cat struct {
	Animal
}

func (animal Animal) MakeSound() {
	fmt.Println(animal.Name, "says:", animal.Sound)
}
func (dog Dog) MakeSound() {
	fmt.Println(dog.Name, "says \"Woof\" 🐶")
}
func (cat Cat) MakeSound() {
	fmt.Println(cat.Name, "says \"Meow\" 🐱")
}

type SoundMaker interface {
	MakeSound()
}

func PlaySound(s SoundMaker) {
	s.MakeSound()
}
func main() {
	dog := Dog{
		Animal: Animal{
			Name:  "Buddy",
			Sound: "???", // не используется, потому что метод переопределён
		},
	}

	cat := Cat{
		Animal: Animal{
			Name:  "Luna",
			Sound: "???",
		},
	}

	PlaySound(dog)
	PlaySound(cat)
}
