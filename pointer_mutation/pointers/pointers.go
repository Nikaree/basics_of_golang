package pointers

import (
	"errors"
	"fmt"
)

var (
	ErrNilPointerMutate = errors.New("nil pointer provided to Mutate")
	ErrNilPointerString = errors.New("nil pointer provided to ReverseString")
)

// Mutate увеличивает значение по указателю на 10
func Mutate(ptr *int) error {
	if ptr == nil {
		return fmt.Errorf("nil pointer provided to Mutate, %w", ErrNilPointerMutate)
	}
	*ptr += 10
	return nil
}

// ReverseString переворачивает строку по указателю
func ReverseString(str *string) error {
	if str == nil {
		return fmt.Errorf("nil pointer provided to ReverseString, %w", ErrNilPointerString)
	}
	runes := []rune(*str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*str = string(runes)
	return nil
}
