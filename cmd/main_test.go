package main

import (
	"testing"
	"time"
)

func TestBasicCalculations(t *testing.T) {
	purchasePrice = 100000
	downpaymentPercentage = 0.10
	feePercentage = 0.20
	monthlyInstallments = 12

	basicCalculations()

	if downpayment != 10000 {
		t.Errorf("Expected downpayment 10000, got %v", downpayment)
	}

	if fee != 20000 {
		t.Errorf("Expected fee 20000, got %v", fee)
	}

	expectedTotal := 110000.0
	if totalRepayable != expectedTotal {
		t.Errorf("Expected total repayable %v, got %v", expectedTotal, totalRepayable)
	}
}

func TestInstallmentCalculation(t *testing.T) {
	purchasePrice = 100000
	downpaymentPercentage = 0.10
	feePercentage = 0.20
	monthlyInstallments = 12
	purchaseDate = time.Date(2026, time.August, 17, 0, 0, 0, 0, time.UTC)

	basicCalculations()
	schedule = nil
	installmentsCalculation()

	if len(schedule) != 12 {
		t.Errorf("Expected 12 installments, got %v", len(schedule))
	}

	expectedInstallment := 9166.67
	if schedule[0].Amount != expectedInstallment {
		t.Errorf("Expected first installment %v, got %v", expectedInstallment, schedule[0].Amount)
	}
}

func TestMonthEndDateHandling(t *testing.T) {
	tests := []struct {
		year     int
		month    time.Month
		expected int
	}{
		{2026, time.February, 28},
		{2024, time.February, 29},
		{2026, time.January, 31},
		{2026, time.April, 30},
	}

	for _, test := range tests {
		result := getLastDayOfMonth(test.year, test.month)
		if result != test.expected {
			t.Errorf("For %v/%v: expected %v days, got %v", test.year, test.month, test.expected, result)
		}
	}
}

func TestDueDateCalculation(t *testing.T) {
	purchasePrice = 100000
	downpaymentPercentage = 0.10
	feePercentage = 0.20
	monthlyInstallments = 12
	purchaseDate = time.Date(2026, time.August, 17, 0, 0, 0, 0, time.UTC)

	basicCalculations()
	schedule = nil
	installmentsCalculation()

	if schedule[0].DueDate.Month() != time.September {
		t.Errorf("Expected first due date in September, got %v", schedule[0].DueDate.Month())
	}

	if schedule[0].DueDate.Day() != 17 {
		t.Errorf("Expected first due date on 17th, got %v", schedule[0].DueDate.Day())
	}
}

func TestDownpaymentCalculation(t *testing.T) {
	tests := []struct {
		price      float64
		percentage float64
		expected   float64
	}{
		{100000, 0.10, 10000},
		{50000, 0.20, 10000},
		{200000, 0.25, 50000},
	}

	for _, test := range tests {
		purchasePrice = test.price
		downpaymentPercentage = test.percentage
		basicCalculations()

		if downpayment != test.expected {
			t.Errorf("For price %v at %v%%: expected %v, got %v", test.price, test.percentage*100, test.expected, downpayment)
		}
	}
}

func TestFeeCalculation(t *testing.T) {
	tests := []struct {
		price      float64
		percentage float64
		expected   float64
	}{
		{100000, 0.20, 20000},
		{50000, 0.10, 5000},
		{200000, 0.15, 30000},
	}

	for _, test := range tests {
		purchasePrice = test.price
		feePercentage = test.percentage
		basicCalculations()

		if fee != test.expected {
			t.Errorf("For price %v at %v%% fee: expected %v, got %v", test.price, test.percentage*100, test.expected, fee)
		}
	}
}
