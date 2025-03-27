package main

import (
	"fmt"
	"os"
	"strconv"
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
)

var (
	fmtIndentMonth = fmt.Sprintf("%%%ds ", len(indent)-1)
	fmtIndentYear  = fmt.Sprintf("%%%dd ", len(indent)-1)
)

func main() {
	printCal(getArgs())
}

func getArgs() (start time.Time, monthCount, lineSkipCount int) {
	// TODO: consider capturing start date from args.
	start = time.Now()
	// show about one quarter by default
	monthCount = 3
	var err error
	args := os.Args[1:]
	if len(args) > 0 {
		monthCount, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("monthCount %q must parse as an integer.\n", args[0])
			os.Exit(1)
		}
		args = args[1:]
	}
	// No line skips by default
	if len(args) > 0 {
		lineSkipCount, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("lineSkipCount %q must parse as an integer.\n", args[0])
			os.Exit(1)
		}
		args = args[1:]
	}
	return
}

func printCal(today time.Time, monthCount, lineSkipCount int) {
	curDay := today

	// dayCount is fuzzy - it's adjusted so that this prints a list
	// of complete weeks (all days filled in from Su to Sa).
	dayCount := monthCount * 31

	// Roll back the date to begin on previous Sunday of current week.
	for curDay.Weekday() != time.Sunday {
		curDay = curDay.Add(-oneDay)
		dayCount--
	}

	// Add enough days to complete a week.
	for dayCount%7 != 0 {
		dayCount++
	}
	// Add two more weeks just because.
	// TODO: assure that the last line represents the end of a month.
	dayCount += 14

	prevYear := -1
	prevMonth := curDay.Add(-oneMonth).Month()
	atLineStart := true

	for n := 0; n < dayCount; n++ {
		if atLineStart {
			if curDay.Year() != prevYear {
				printYearHeader(curDay)
				prevYear = curDay.Year()
			}
			if month := curDay.Add(oneWeek - oneDay).Month(); month != prevMonth {
				printMonthIndent(month)
				prevMonth = month
			} else {
				fmt.Print(indent)
			}
			atLineStart = false
		}
		if curDay == today {
			fmt.Printf(fmtToday, curDay.Day())
		} else {
			fmt.Printf(fmtNotToday, curDay.Day())
		}
		if curDay.Weekday() == time.Saturday {
			fmt.Println()
			atLineStart = true
			for i := 0; i < lineSkipCount; i++ {
				fmt.Println()
			}
		}
		curDay = curDay.Add(oneDay)
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
