package flags

import (
	"github.com/pborman/getopt"
)

type Flags struct {
	FlagK    int
	FlagN    bool
	FlagR    bool
	FlagU    bool
	FlagB    bool
	FlagM    bool
	FlagC    bool
	FileName string
}

func Parse() *Flags {
	flagK := getopt.Int('k', 0, "for setting start column")
	flagN := getopt.Bool('n', "for number sort")
	flagR := getopt.Bool('r', "for reversed sort")
	flagU := getopt.Bool('u', "for not doubling lines")
	flagB := getopt.Bool('b', "for removing trailing blanks")
	flagM := getopt.Bool('M', "for sorting by month")
	flagC := getopt.Bool('c', "check if input is sorted")
	getopt.Parse()

	flags := Flags{
		FlagK: *flagK,
		FlagN: *flagN,
		FlagR: *flagR,
		FlagU: *flagU,
		FlagB: *flagB,
		FlagM: *flagM,
		FlagC: *flagC,
	}

	args := getopt.Args()
	if len(args) > 0 {
		flags.FileName = args[0]
	}

	return &flags
}
