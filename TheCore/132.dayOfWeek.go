package main

import "time"

func dayOfWeek(birthdayDate string) int {
	layout := "01-02-2006"
	date, _ := time.Parse(layout, birthdayDate)
	weekday := date.Weekday()
	day := date.Day()

	for i := 1; ; i++ {
		tmpDate := date.AddDate(i, 0, 0)
		if tmpDate.Weekday() == weekday && tmpDate.Day() == day {
			return i
		}
	}

	return 0
}
