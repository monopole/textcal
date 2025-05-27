package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	oneDay       = 24 * time.Hour
	oneWeek      = 7 * oneDay
	oneMonth     = 4 * oneWeek
	indent       = "        "
	fmtDayHeader = " %2s "
	fmtToday     = "[%2d]"
	fmtNotToday  = " %2d "

	// Arbitrary limits to prevent huge output
	maxMonthCount    = 24
	maxLineSkipCount = 5

	usage = `
  textcal {monthCount} {lineSkipCount}

    monthCount defaults to 3
    lineSkipCount defaults to 0

`
)

var (
	fmtIndentMonth = fmt.Sprintf("%%%ds ", len(indent)-1)
	fmtIndentYear  = fmt.Sprintf("%%%dd ", len(indent)-1)
)

func main() {
	printCal(getArgs())
}

func getArgs() (monthCount, lineSkipCount int) {
	// show about one quarter by default
	monthCount = 3
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	if strings.Contains(args[0], "-h") {
		fmt.Print(usage)
		os.Exit(0)
	}

	var err error
	monthCount, err = strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("monthCount %q must parse as an integer.\n", args[0])
		os.Exit(1)
	}
	if monthCount > maxMonthCount {
		fmt.Printf("monthCount %d exceeds max of %d\n",
			monthCount, maxMonthCount)
		os.Exit(1)
	}

	args = args[1:]
	if len(args) == 0 {
		return
	}

	// No line skips by default
	lineSkipCount, err = strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("lineSkipCount %q must parse as an integer.\n", args[0])
		os.Exit(1)
	}
	if lineSkipCount > maxLineSkipCount {
		fmt.Printf("lineSkipCount %d exceeds max of %d\n",
			lineSkipCount, maxLineSkipCount)
		os.Exit(1)
	}
	return
}

func printCal(monthCount, lineSkipCount int) {
	today := time.Now()

	dayPtr := today
	// Roll back dayPtr to start the calendar on the Sunday of previous week.
	// Usually one wants to know what the previous week looked like.
	dayPtr = dayPtr.Add(-7 * oneDay)
	for dayPtr.Weekday() != time.Sunday {
		dayPtr = dayPtr.Add(-oneDay)
	}

	// dayCount is fuzzy count of how many days to show in the calendar.
	// It gets adjusted to assure that weeks are completed from Su to Sa.
	dayCount := monthCount * 30
	// Add enough days to complete a week.
	for dayCount%7 != 0 {
		dayCount++
	}
	// Add two weeks to compensate for the up to two weeks we went
	// backward with dayPtr.
	dayCount += 14

	prevYear := -1
	prevMonth := dayPtr.Add(-oneMonth).Month()
	atLineStart := true

	for n := 0; n < dayCount; n++ {
		if atLineStart {
			// Handle indentation.
			if dayPtr.Year() != prevYear {
				printYearHeader(dayPtr)
				prevYear = dayPtr.Year()
			}
			if month := dayPtr.Add(oneWeek - oneDay).Month(); month != prevMonth {
				printMonthIndent(month)
				prevMonth = month
			} else {
				fmt.Print(indent)
			}
			atLineStart = false
		}
		if dayPtr == today {
			fmt.Printf(fmtToday, dayPtr.Day())
		} else {
			fmt.Printf(fmtNotToday, dayPtr.Day())
		}
		if dayPtr.Weekday() == time.Saturday {
			fmt.Println()
			atLineStart = true
			for i := 0; i < lineSkipCount; i++ {
				fmt.Println()
			}
		}
		dayPtr = dayPtr.Add(oneDay)
	}
	if !atLineStart {
		fmt.Println()
	}
}

func printMonthIndent(month time.Month) {
	fmt.Printf(fmtIndentMonth, month.String()[0:3])
}

func printYearHeader(date time.Time) {
	fmt.Printf(fmtIndentYear, date.Year())
	for _, d := range []time.Weekday{
		time.Sunday, time.Monday, time.Tuesday,
		time.Wednesday, time.Thursday, time.Friday, time.Saturday} {
		fmt.Printf(fmtDayHeader, d.String()[:2])
	}
	fmt.Println()
}
