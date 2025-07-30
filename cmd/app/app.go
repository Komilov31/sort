package app

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Komilov31/sort/internal/flags"
	"github.com/Komilov31/sort/internal/sort"
)

type App struct {
	sort  *sort.Sort
	flags *flags.Flags
}

func New(sort *sort.Sort, flags *flags.Flags) *App {
	return &App{
		sort:  sort,
		flags: flags,
	}
}

func (a *App) Run(output io.Writer) {
	source := a.determineSource()
	defer func() {
		if err := source.Close(); err != nil {
			log.Fatalf("could not close file: %v\n", err)
		}
	}()

	a.sort.ProcessInput(source, a.flags.FlagB)

	if a.flags.FlagU {
		a.sort.UniqueStrings()
	}

	if !a.flags.FlagC {
		a.sort.Sort()
	}

	if a.flags.FlagK != 0 || a.flags.FlagN {
		a.sort.SortByColumn(a.flags.FlagK, a.flags.FlagN)
	}

	if a.flags.FlagM {
		a.sort.SortByMonth(a.flags.FlagK)
	}

	if a.flags.FlagR {
		a.sort.ReversedSort()
	}

	if a.flags.FlagC {
		a.sort.CheckIfSorted(a.flags.FileName)
		return
	}

	a.sort.PrintResult(output)
}

// функция для определения то, из файла читать или из STDIN
func (a *App) determineSource() *os.File {
	source := os.Stdin
	if a.flags.FileName != "" {
		fileName := a.flags.FileName

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("sort: cannot read: %s: No such file or directory", fileName)
			os.Exit(1)
		}
		source = file
	}
	return source
}
