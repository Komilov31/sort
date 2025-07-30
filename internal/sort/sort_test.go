package sort

import (
	"slices"
	"testing"
)

var sortInstance = New()

func TestUniqueStrings(t *testing.T) {
	testCases := [][]string{
		{"apple", "яблоко", "green", "walk", "яблоко", "green"},
		{"яблоко", "груша", "груша", "ходить"},
		{"22 22", "3, 35", "22 22", "3, 35", "1", "2"},
		{"sdka we wsd", "sdka we wsd", "3 apple", "4 grape"},
	}

	for _, testCase := range testCases {
		sortInstance.storage = testCase
		mapOfStorage := make(map[string]int)
		sortInstance.UniqueStrings()

		for _, value := range sortInstance.storage {
			mapOfStorage[value]++
		}

		for _, value := range mapOfStorage {
			if value != 1 {
				t.Error("found duplicate element")
			}
		}
	}

}

func TestSortByColumn(t *testing.T) {
	//проверка когда столбец начало
	sortInstance.storage = []string{
		"Lucke Rockhold", "Daniel Cormier", "Jon Jones", "Conor Macgregor", "\n",
	}
	expectedResult := []string{"\n", "Conor Macgregor", "Daniel Cormier", "Jon Jones", "Lucke Rockhold"}
	sortInstance.SortByColumn(1, false)
	if !slices.Equal(sortInstance.storage, expectedResult) {
		t.Errorf("not correct order")
	}

	// проверка когда столбец > 1
	sortInstance.storage = []string{
		"qwerty qwerty", "hello world", "добрый вечер", "911",
	}
	expectedResult = []string{"911", "qwerty qwerty", "hello world", "добрый вечер"}
	sortInstance.SortByColumn(2, false)
	if !slices.Equal(sortInstance.storage, expectedResult) {
		t.Errorf("not correct order")
	}

	// проверка по числовым значениям
	sortInstance.storage = []string{
		"30 виноград", "10 яблок", "911", "20 груш", "231",
	}
	expectedResult = []string{"10 яблок", "20 груш", "30 виноград", "231", "911"}
	sortInstance.SortByColumn(1, true)
	if !slices.Equal(sortInstance.storage, expectedResult) {
		t.Errorf("not correct order")
	}

	// проверка по числовым значениям по столбцам n > 1
	sortInstance.storage = []string{
		"24 21 10 apples", "22 me 5", "13 siya 30 long", "2sakdjf sl 155",
	}
	expectedResult = []string{"22 me 5", "24 21 10 apples", "13 siya 30 long", "2sakdjf sl 155"}
	sortInstance.SortByColumn(3, true)
	if !slices.Equal(sortInstance.storage, expectedResult) {
		t.Errorf("not correct order")
	}
}

func TestSortByMonth(t *testing.T) {
	//проверка когда столбец начало
	sortInstance.storage = []string{
		"happened in Mar", "Sep is nice month", "I love Jan", "Aug is very hot", "Lucke Rockhold",
	}
	expectedResult := []string{"happened in Mar", "I love Jan", "Lucke Rockhold", "Aug is very hot", "Sep is nice month"}
	sortInstance.SortByMonth(1)
	if !slices.Equal(sortInstance.storage, expectedResult) {
		t.Errorf("not correct order")
	}

	// проверка когда столбец > 1
	sortInstance.storage = []string{
		"qwerty qwerty Nov", "hello world Oct", "добрый вечер Mar", "911 32 Jan",
	}
	expectedResult = []string{"911 32 Jan", "добрый вечер Mar", "hello world Oct", "qwerty qwerty Nov"}
	sortInstance.SortByMonth(3)
	if !slices.Equal(sortInstance.storage, expectedResult) {
		t.Errorf("not correct order")
	}
}
