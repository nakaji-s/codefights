package main

import "time"

func regularMonths(currMonth string) string {
	layout := "01-2006"
	date, _ := time.Parse(layout, currMonth)

	for i := 1; ; i++ {
		tmpDate := date.AddDate(0, i, 0)
		if tmpDate.Weekday() == time.Monday {
			return tmpDate.Format(layout)
		}
	}

	return ""
}
