package main

import (
	"basics/json/safeio"
	"fmt"
	"os"
)

func main() {
	// Проверка WriteFileAtomic
	if err := safeio.WriteFileAtomic("greet.txt", []byte("Hello!"), 0644); err != nil {
		fmt.Println("write failed:", err)
	}

	// Проверка WithFile с искусственной паникой
	err := safeio.WithFile("data.txt", os.O_CREATE|os.O_RDWR, 0644,
		func(f *os.File) error {
			_, err := f.WriteString("test")
			if err != nil {
				return err
			}
			panic("disk full")
		})
	if err != nil {
		fmt.Println("operation failed:", err)
	}
}
