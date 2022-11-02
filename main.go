package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tsuen4/onemonth/pkg/during"
)

var (
	year       int
	month      int
	timeLayout string
)

func init() {
	var (
		usageYear   = "number of years"
		defaultYear = time.Now().Local().Year()
	)
	flag.IntVar(&year, "year", defaultYear, usageYear)
	flag.IntVar(&year, "y", defaultYear, usageYear+" (shorthand)")

	var (
		usageMonth   = "number of months"
		defaultMonth = int(time.Now().Local().Month())
	)
	flag.IntVar(&month, "month", defaultMonth, usageMonth)
	flag.IntVar(&month, "m", defaultMonth, usageMonth+" (shorthand)")

	var (
		usageTimeLayout   = "string of time format layout."
		defaultTimeLayout = "2006/01/02: "
	)
	flag.StringVar(&timeLayout, "layout", defaultTimeLayout, usageTimeLayout)
	flag.StringVar(&timeLayout, "l", defaultTimeLayout, usageTimeLayout+" (shorthand)")
}

func main() {
	flag.Parse()

	month, err := during.NewOneMonth(year, month)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	month.Iterate(func(day time.Time) {
		fmt.Println(day.Format(timeLayout))
	})
}
