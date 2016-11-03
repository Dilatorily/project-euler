package main

import (
	"testing"
	"strconv"
)

var isMultipleTestingTable = []struct {
	number int
	radix int
	isMultiple bool
	errorMessage string
}{
	{12, 1, true, "Failed to detect that 12 is a multiple of 1"},
	{12, 2, true, "Failed to detect that 12 is a multiple of 2"},
	{12, 3, true, "Failed to detect that 12 is a multiple of 3"},
	{12, 4, true, "Failed to detect that 12 is a multiple of 4"},
	{12, 5, false, "Failed to detect that 12 is not a multiple of 5"},
	{12, 6, true, "Failed to detect that 12 is a multiple of 6"},
	{12, 12, true, "Failed to detect that 12 is a multiple of 12"},
}

func TestIsMultiple(t *testing.T) {
	for _, entry := range isMultipleTestingTable {
		isMultiple := IsMultiple(entry.number, entry.radix)
		if isMultiple != entry.isMultiple {
			t.Error(entry.errorMessage)
		}
	}
}

var isMultipleOf3Or5TestingTable = []struct {
	number int
	isMultiple bool
	errorMessage string
}{
	{1, false, "Failed to detect that 1 is not a multiple of 3 or 5"},
	{2, false, "Failed to detect that 2 is not a multiple of 3 or 5"},
	{3, true, "Failed to detect that 3 is a multiple of 3 or 5"},
	{4, false, "Failed to detect that 4 is not a multiple of 3 or 5"},
	{5, true, "Failed to detect that 5 is not a multiple of 3 or 5"},
	{6, true, "Failed to detect that 6 is not a multiple of 3 or 5"},
	{7, false, "Failed to detect that 7 is not not a multiple of 3 or 5"},
	{8, false, "Failed to detect that 8 is not not a multiple of 3 or 5"},
	{15, true, "Failed to detect that 15 is a multiple of 3 or 5"},
	{18, true, "Failed to detect that 18 is a multiple of 3 or 5"},
	{19, false, "Failed to detect that 19 is not a multiple of 3 or 5"},
	{20, true, "Failed to detect that 20 is a multiple of 3 or 5"},
	{21, true, "Failed to detect that 21 is not a multiple of 3 or 5"},
	{25, true, "Failed to detect that 25 is a multiple of 3 or 5"},
	{30, true, "Failed to detect that 30 is a multiple of 3 or 5"},
}

func TestIsMultipleOf3Or5(t *testing.T) {
	for _, entry := range isMultipleOf3Or5TestingTable {
		isMultipleOf3Or5 := IsMultipleOf3Or5(entry.number)
		if isMultipleOf3Or5 != entry.isMultiple {
			t.Error(entry.errorMessage)
		}
	}
}

var getSumOfMultiplesTestingTable = []struct {
	limit int
	sum int
}{
	{1, 0},
	{2, 0},
	{3, 0},
	{4, 3},
	{5, 3},
	{6, 8},
	{10, 23},
}

func TestGetSumOfMultiples(t *testing.T) {
	for _, entry := range getSumOfMultiplesTestingTable {
		sum := GetSumOfMultiples(entry.limit)
		if sum != entry.sum {
			t.Error("The sum of multiples of 3 or 5 of all natural numbers below " + strconv.Itoa(entry.limit) + " is not " + strconv.Itoa(entry.sum))
		}
	}
}
