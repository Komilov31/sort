package sort

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Sort struct {
	storage []string
}

func New() *Sort {
	return &Sort{
		storage: []string{},
	}
}

func (s *Sort) Sort() {
	sort.Slice(s.storage, func(i, j int) bool {
		return strings.ToLower(s.storage[i]) < strings.ToLower(s.storage[j])
	})
}

func (s *Sort) ReversedSort() {
	slices.Reverse(s.storage)
}

// функция отбирает уникальные строки из storage
func (s *Sort) UniqueStrings() {
	storageUnique := make(map[string]struct{})

	for _, line := range s.storage {
		storageUnique[line] = struct{}{}
	}

	storage := make([]string, 0, len(s.storage))
	for k := range storageUnique {
		storage = append(storage, k)
	}

	s.storage = storage
}

// функция сортирует по определенной колонне(столбец), используя функцию из stdlib
// если не указан аргумент сортировки по определенному столбцу,
// то сортируется по первому столбцу
// также если параметр numeric будет true, то сортировка будет по числовому зн.
func (s *Sort) SortByColumn(column int, numeric bool) {
	position := column
	if column != 0 {
		position = column - 1
	}
	sort.SliceStable(s.storage, func(i, j int) bool {
		fieldsI := strings.Fields(strings.ToLower(s.storage[i]))
		fieldsJ := strings.Fields(strings.ToLower(s.storage[j]))

		// если у левой строки нету нужного столбца, то он по умолчанию меншье
		if len(fieldsI) <= position {
			return true
		}

		if len(fieldsJ) <= position {
			return false
		}

		// если указан флаг n(сортировка по числу), то сортируем по числам
		// нужного столбца
		if numeric {
			fieldsINum, _ := strconv.Atoi(fieldsI[position])
			fieldsJNum, _ := strconv.Atoi(fieldsJ[position])
			return fieldsINum < fieldsJNum
		}

		return fieldsI[position] < fieldsJ[position]
	})

	pos := len(s.storage) - 1
	for i, value := range s.storage {
		fields := strings.Fields(value)
		if len(fields) >= column {
			pos = i
			break
		}
	}

	sort.Slice(s.storage[:pos], func(i, j int) bool {
		return strings.ToLower(s.storage[i]) < strings.ToLower(s.storage[j])
	})
}

func (s *Sort) CheckIfSorted(fileName string) {
	if fileName == "" {
		fileName = "-"
	}

	for i, j := 0, 1; j < len(s.storage); i, j = i+1, j+1 {
		if s.storage[j] < s.storage[i] {
			fmt.Printf("sort: %s:%d: disorder: %s", fileName, j+1, s.storage[j])
			return
		}
	}
}

func (s *Sort) SortByMonth(column int) {
	if column > 0 {
		column = column - 1
	}
	monthOrder := map[string]int{
		"Jan": 0, "Feb": 1, "Mar": 2, "Apr": 3, "May": 4, "Jun": 5,
		"Jul": 6, "Aug": 7, "Sep": 8, "Oct": 9, "Nov": 10, "Dec": 11,
	}

	sort.Slice(s.storage, func(i, j int) bool {
		// Найти месяц в строке, путем проверки каждого столбца с содержанием мапы
		findMonth := func(input string) int {
			columns := strings.Fields(input)
			//если в строке нет нужной колонки, то она меньше
			if len(columns) <= column {
				return -1
			}
			if order, ok := monthOrder[columns[column]]; ok {
				return order
			}

			return -1
		}

		return findMonth(s.storage[i]) < findMonth(s.storage[j])
	})

	pos := len(s.storage) - 1
	for i, value := range s.storage {
		fields := strings.Fields(value)
		if len(fields) > 0 {
			if _, ok := monthOrder[fields[column]]; ok {
				pos = i
				break
			}
		}
	}

	sort.Slice(s.storage[:pos], func(i, j int) bool {
		return strings.ToLower(s.storage[i]) < strings.ToLower(s.storage[j])
	})
}

// чтение инпута и сохранение в storage
func (s *Sort) ProcessInput(input io.Reader, deleteTrailingBlank bool) {
	reader := bufio.NewReader(input)
	for {
		nextString, err := reader.ReadString('\n')
		if err == io.EOF {
			nextString = nextString + "\n"
		}
		if deleteTrailingBlank {
			nextString = strings.TrimRight(nextString, "\t\r")
		}
		s.storage = append(s.storage, nextString)

		if err == io.EOF {
			break
		}
	}
}

func (s *Sort) PrintResult(output io.Writer) {
	for _, value := range s.storage {
		_, err := output.Write([]byte(value))
		if err != nil {
			log.Fatalf("could not write result to output: %v\n", err)
		}
	}
}
