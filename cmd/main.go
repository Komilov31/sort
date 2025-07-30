package main

import (
	"os"

	"github.com/Komilov31/sort/cmd/app"
	"github.com/Komilov31/sort/internal/flags"
	"github.com/Komilov31/sort/internal/sort"
)

func main() {
	flags := flags.Parse()
	sort := sort.New()
	app := app.New(sort, flags)

	app.Run(os.Stdout)
}
