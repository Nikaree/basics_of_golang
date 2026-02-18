package crm

import (
	"fmt"
	"time"
)

// Entity Базовая «запись в БД»
type Entity struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Person Персона (сотрудник, контакт…)
type Person struct {
	Entity
	Name  string
	Email string
}

// Company Компания-клиент
type Company struct {
	Entity
	Name    string
	Address Address
}

// Address Почтовый адрес
type Address struct {
	Street  string
	City    string
	Zip     string
	Country string
}

// Update обновляет поле UpdatedAt текущим временем.
func (e *Entity) Update() {
	e.UpdatedAt = time.Now()
}

// Greet возвращает приветствие с именем и почтой.
// Если Email пустой, выводите пустые скобки ().
func (p Person) Greet() string {
	email := fmt.Sprintf("(%s)", p.Email)
	greeting := fmt.Sprintf("Hi, I'm %s! %s", p.Name, email)
	return greeting
}

// FullAddress возвращает адрес в виде "Street, City, Zip, Country".
func (a Address) FullAddress() string {
	address := fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.Zip, a.Country)
	return address
}

// Location делегирует метод FullAddress.
func (c Company) Location() string {
	loc := c.Address.FullAddress()
	return loc
}

// Contact возвращает "<name> — <address.fulladdress()>".
func (c Company) Contact() string {
	contact := fmt.Sprintf("<%s> — <%s>", c.Name, c.Location())
	return contact
}
