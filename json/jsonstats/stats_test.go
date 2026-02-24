package jsonstats

import (
	"strings"
	"testing"
)

func TestCalculateStats_Success(t *testing.T) {
	jsonInput := `
	[
		{"user":"u1","category":"books","spent":100},
		{"user":"u2","category":"tech","spent":500},
		{"user":"u1","category":"books","spent":200}
	]
	`

	report, err := CalculateStats(strings.NewReader(jsonInput))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if report.RecordsTotal != 3 {
		t.Errorf("RecordsTotal = %d, want 3", report.RecordsTotal)
	}

	if report.UniqueUsers != 2 {
		t.Errorf("UniqueUsers = %d, want 2", report.UniqueUsers)
	}

	if report.SumSpent != 800 {
		t.Errorf("SumSpent = %d, want 800", report.SumSpent)
	}

	// Максимальные затраты у категории tech (500)
	if report.TopCategory.Category != "tech" {
		t.Errorf("TopCategory.Category = %s, want tech", report.TopCategory.Category)
	}

	if report.TopCategory.Spent != 500 {
		t.Errorf("TopCategory.Spent = %d, want 500", report.TopCategory.Spent)
	}
}
