package main

import (
	// Укажите правильный путь к вашему модулю
	"basics/interfaces/notification_system/notification"
	"fmt" // Может понадобиться для вывода ошибок
)

func main() {
	// Создаем базовые уведомители
	console := notification.ConsoleNotifier{}
	email := notification.EmailNotifier{} // Предположим, он иногда возвращает ошибку для теста Retry

	// Создаем декорированные уведомители
	// 1. Консольный уведомитель с форматированием
	formattedConsole := notification.FormattedNotifier{Notifier: console}

	// 2. Email уведомитель с 3 попытками отправки
	retryEmail := notification.RetryNotifier{Notifier: email, MaxAttempts: 3}

	// Используем их
	fmt.Println("Отправка через форматированную консоль:")
	err := formattedConsole.Send("Привет, это тестовое сообщение!")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println("\nОтправка Email с повторами:")
	err = retryEmail.Send("Важное уведомление")
	if err != nil {
		fmt.Println("Итоговая ошибка Email:", err)
	}

	// Можно комбинировать декораторы!
	fmt.Println("\nОтправка через консоль с форматированием и повторами:")
	retryFormattedConsole := notification.RetryNotifier{
		Notifier:    notification.FormattedNotifier{Notifier: console},
		MaxAttempts: 2,
	}
	err = retryFormattedConsole.Send("Сообщение с двойным декорированием")
	if err != nil {
		fmt.Println("Итоговая ошибка комбинированного:", err)
	}
}
