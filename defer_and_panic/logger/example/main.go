package main

import (
	"basics/defer_and_panic/logger"
)

func main() {
	// Устанавливаем минимальный уровень логирования
	logger.SetMinLevel(logger.INFO)

	// DEBUG сообщение не будет выведено (ниже минимального уровня)
	logger.Debug("This is debug")

	// Эти сообщения будут выведены
	logger.Info("Server started")
	logger.Warning("Low memory")
	logger.Error("Connection failed")

	// Обработка паники
	defer func() {
		if r := recover(); r != nil {
			logger.Info("Recovered from panic:", r)
		}
	}()

	// Критическая ошибка с дополнительной информацией
	logger.Fatal("Fatal error:", "userID=", 42, "err=", "db down")

	// Эта строка не выполнится из-за panic
	logger.Info("This won't be printed")
}
