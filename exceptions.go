package main

import "time"

// eventaully, this should be read from a file, but for now this will work
func loadExceptions() []string {
	exceptions := []string{
		"2018-09-03 Labour Day",
		"2018-10-05 Teacher Conf",
		"2018-10-08 Thanksgiving",
		"2018-10-26 No school",
		"2018-11-12 In lieu of Rememberance Day",
		"2018-12-07 No School",
		"2018-12-24 Vacation",
		"2018-12-25 Vacation",
		"2018-12-26 Vacation",
		"2018-12-27 Vacation",
		"2018-12-28 Vacation",
		"2018-12-31 Vacation",
		"2019-01-01 Vacation",
		"2019-01-02 Vacation",
		"2019-01-03 Vacation",
		"2019-01-04 Vacation",
		"2019-01-25 no school",
		"2019-02-18 Vacation",
		"2019-02-19 Vacation",
		"2019-02-20 Vacation",
		"2019-02-21 Vacation",
		"2019-02-22 Vacation",
		"2019-03-22 Vacation",
		"2019-04-19 good Friday",
		"2019-04-22 Vacation",
		"2019-04-23 Vacation",
		"2019-04-24 Vacation",
		"2019-04-25 Vacation",
		"2019-04-26 Vacation",
		"2019-05-17 no school",
		"2019-05-20 no school",
		"2019-06-14 No school",
		"2019-06-28 No More School",
		"2018-02-16 Teacher confrence",
	}
	return exceptions
}

var dayFormat = "2006-01-02"
var firstDay, _ = time.Parse("2006-01-02 15:04", "2018-09-03 08:00")
var lastDay, _ = time.Parse("2006-01-02 15:04", "2019-06-28 08:00")
var numberOfDays = 5
