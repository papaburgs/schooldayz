// Package schooldayz is a little app to generate a slice of days, meant to be output
// into google calendar via a csv export
// buffalo or other exporters to come later
package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Day is a struct that holds information about that day, for exporting purposes
type Day struct {
	date        time.Time
	subject     string
	allDayEvent bool
	description string
	noSchoolDay bool
}

func (d Day) GoogleCalDay() string {
	s := fmt.Sprintf("%s,%s,True,Added by script",
		d.subject,
		d.date.Format("01/02/06"),
	)
	return s
}

func (d Day) getDayString() string {
	return d.date.Format(dayFormat)
}

func NewDay(day, subject, description string, allDay, noSchool bool) (Day, error) {
	var err error
	d := Day{
		subject:     subject,
		description: description,
		allDayEvent: allDay,
		noSchoolDay: noSchool,
	}
	d.date, err = time.Parse(dayFormat, day)
	if err != nil {
		return d, err
	}
	return d, nil
}

func nextDay(currentDay int) int {
	if currentDay == numberOfDays {
		return 1
	}
	return currentDay + 1
}

func isWeekend(d time.Time) bool {
	if d.Weekday() == time.Sunday || d.Weekday() == time.Saturday {
		return true
	}
	return false

}

func buildDaysList(exceptions []Day) ([]Day, error) {
	currentDay := firstDay
	twentyFourHours, _ := time.ParseDuration("24h")
	var isException bool
	var thisDay Day
	daysList := []Day{}
	var dayIndex = 1

	// loop while the last day is still after the current iteration
	// if its a weekend, go to next day
	// if its an exception,
	//    add that day into the list, go to next
	// else add a new day in with a subject of the day number
	//      increase the day number
	for lastDay.After(currentDay) {
		if !isWeekend(currentDay) {
			isException = false
			for _, e := range exceptions {
				if currentDay.Format(dayFormat) == e.getDayString() {
					isException = true
					thisDay = e
					break
				}
			}
			if !isException {
				thisDay, _ = NewDay(currentDay.Format(dayFormat),
					fmt.Sprintf("Day %v", dayIndex),
					"Added by script",
					true,
					false,
				)
				dayIndex = nextDay(dayIndex)
			}
			daysList = append(daysList, thisDay)

		}
		// always increment day
		currentDay = currentDay.Add(twentyFourHours)

	}
	return daysList, nil
}

func main() {
	var (
		s      []string
		nd     Day
		err    error
		ex     []Day
		ex_raw []string
		days   []Day
	)

	ex_raw = loadExceptions()
	ex = []Day{}
	for _, x := range ex_raw {
		s = strings.SplitN(x, " ", 2)
		nd, err = NewDay(s[0], s[1], " ", true, true)
		if err != nil {
			log.Println(err)
			continue
		}
		ex = append(ex, nd)
	}
	days, _ = buildDaysList(ex)

	fmt.Println("Subject,Start Date,All Day Event, Description")
	for _, x := range days {
		fmt.Println(x.GoogleCalDay())
	}

}
