package main

import (
	"basics/data_structure/embedded_record/crm"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	person := crm.Person{
		Entity: crm.Entity{ID: 7, CreatedAt: now, UpdatedAt: now},
		Name:   "Alice",
		Email:  "alice@example.com",
	}

	company := crm.Company{
		Entity: crm.Entity{ID: 42, CreatedAt: now},
		Name:   "ACME Inc.",
		Address: crm.Address{
			Street:  "123 Main St",
			City:    "Metropolis",
			Zip:     "10001",
			Country: "USA",
		},
	}

	fmt.Println(person.Greet())     // Hi, I'm Alice (alice@example.com)
	fmt.Println(company.Location()) // 123 Main St, Metropolis, 10001, USA
	fmt.Println(company.Contact())  // ACME Inc. — 123 Main St, Metropolis, 10001, USA

	time.Sleep(time.Millisecond)
	company.Update()
	fmt.Println(company.UpdatedAt.After(now)) // true
}
