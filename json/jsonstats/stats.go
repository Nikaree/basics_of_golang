package jsonstats

import (
	"encoding/json"
	"fmt"
	"io"
)

// Report содержит статистику по покупкам.
type Report struct {
	RecordsTotal int           `json:"records_total"` // всего записей
	UniqueUsers  int           `json:"unique_users"`  // разные пользователи
	SumSpent     int64         `json:"sum_spent"`     // сумма всех покупок
	TopCategory  TopByCategory `json:"top_category"`  // топ-категория по сумме
}

// TopByCategory содержит категорию с максимальными затратами.
type TopByCategory struct {
	Category string `json:"category"`
	Spent    int64  `json:"spent"`
}

type purchase struct {
	User     string `json:"user"`
	Category string `json:"category"`
	Spent    int64  `json:"spent"`
}

// CalculateStats читает покупки из r и формирует отчёт.
func CalculateStats(r io.Reader) (Report, error) {
	// создаем декодер
	decoder := json.NewDecoder(r)

	// Читаем первый токен — должен быть '['
	_, err := decoder.Token()
	if err != nil {
		return Report{}, fmt.Errorf("invalid json: %w", err)
	}

	var report Report

	// Подготовка структур для подсчётов
	var (
		totalRecords int
		sumSpent     int64
		users        = make(map[string]struct{}) // set пользователей
		categorySum  = make(map[string]int64)    // сумма по категориям
	)

	for decoder.More() {
		var p purchase

		// Декодируем одну покупку
		if err := decoder.Decode(&p); err != nil {
			return Report{}, fmt.Errorf("decode error: %w", err)
		}

		totalRecords++
		sumSpent += p.Spent
		users[p.User] = struct{}{}
		categorySum[p.Category] += p.Spent

		// поиск топ категории
		var top TopByCategory
		for category, sum := range categorySum {
			if top.Spent < sum {
				top.Spent = sum
				top.Category = category
			}
		}

		// формирует отчет
		report = Report{
			RecordsTotal: totalRecords,
			UniqueUsers:  len(users),
			SumSpent:     sumSpent,
			TopCategory:  top,
		}

	}

	// Читаем закрывающую ']'
	if _, err := decoder.Token(); err != nil {
		return Report{}, fmt.Errorf("invalid json closing: %w", err)
	}

	return report, nil

}
