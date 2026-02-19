package ifacewarmup

import (
	"fmt"
	"io"
)

// PrintAll выводит v.String() каждой строки в w + '\n'.
// После УСПЕШНОЙ записи, если v реализует io.Closer, вызывается Close().
// Первая ошибка (write или close) немедленно возвращается.
func PrintAll(w io.Writer, values ...fmt.Stringer) error {
	if w == nil {
		return nil
	}

	for _, value := range values {
		if value == nil {
			continue
		}

		// Пишем строку + перевод строки
		_, err := io.WriteString(w, value.String()+"\n")
		if err != nil {
			return err
		}

		// После успешной записи проверяем,
		// реализует ли значение io.Closer
		if closer, ok := value.(io.Closer); ok {
			if err := closer.Close(); err != nil {
				return err
			}
		}
	}
	return nil
}
