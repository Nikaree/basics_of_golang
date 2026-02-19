package notification

import (
	"fmt"
	"time"
)

// FormattedNotifier - декоратор, добавляющий форматирование (например, timestamp) к сообщению.
type FormattedNotifier struct {
	// Notifier - вложенный уведомитель, которому будет делегирована отправка.
	Notifier Notifier
}

// Send для FormattedNotifier.
// Он должен отформатировать сообщение (например, добавить текущее время time.Now().Format(time.RFC3339))
// и затем вызвать метод Send вложенного Notifier с отформатированным сообщением.
// Должен возвращать ошибку от вложенного Notifier.
func (f FormattedNotifier) Send(message string) error {
	currentTime := time.Now().Format(time.RFC3339)
	format := message + "\t" + currentTime
	return f.Notifier.Send(format)
}

// RetryNotifier - декоратор, повторяющий попытку отправки сообщения при ошибке.
type RetryNotifier struct {
	// Notifier - вложенный уведомитель.
	Notifier Notifier
	// MaxAttempts - максимальное количество попыток отправки.
	MaxAttempts int
}

// Send для RetryNotifier.
// Он должен пытаться вызвать метод Send вложенного Notifier.
// Если Send возвращает ошибку, нужно подождать (например, time.Sleep(1 * time.Second))
// и повторить попытку, но не более MaxAttempts раз.
// Если хотя бы одна попытка успешна (ошибки нет), метод должен вернуть nil.
// Если все MaxAttempts попыток не удались, метод должен вернуть ошибку (например, используя fmt.Errorf).
func (r RetryNotifier) Send(message string) error {
	var lastError error
	for i := 0; i <= r.MaxAttempts; i++ {
		err := r.Notifier.Send(message)
		lastError = err
		if err == nil {
			return nil
		}
		if i != r.MaxAttempts {
			time.Sleep(1 * time.Second)
		}
	}
	return fmt.Errorf("failed to send message:%w", lastError)
}
