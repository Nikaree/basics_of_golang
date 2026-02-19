package notification

import "fmt"

// Notifier определяет контракт для отправки уведомлений.
type Notifier interface {
	// Send отправляет сообщение. Возвращает ошибку, если отправка не удалась.
	Send(message string) error
}

// ConsoleNotifier реализует Notifier для вывода в консоль.
type ConsoleNotifier struct{}

// Send для ConsoleNotifier
// Он должен выводить сообщение в консоль, например, в формате "Console: {message}".
// Должен возвращать nil в качестве ошибки.
func (c ConsoleNotifier) Send(message string) error {
	fmt.Printf("[CONSOLE]: %s\n ", message)
	return nil
}

// EmailNotifier реализует Notifier, имитируя отправку email.
type EmailNotifier struct{}

// Send для EmailNotifier.
// Он должен имитировать отправку email, например, выводя "Email sent: {message}".
// Должен возвращать nil в качестве ошибки.
func (e EmailNotifier) Send(message string) error {
	fmt.Printf("[EMAIL SENT]: %s\n ", message)
	return nil
}
