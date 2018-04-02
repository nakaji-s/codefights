package main

import "time"

func holiday(x int, weekDay string, month string, yearNumber int) int {
	m := map[string]time.Month{}
	for i := time.January; i < 13; i++ {
		m[i.String()] = i
	}

	date := time.Date(yearNumber, m[month], 1, 0, 0, 0, 0, time.UTC)
	for date.Weekday().String() != weekDay {
		date = date.AddDate(0, 0, 1)
	}
	date = date.AddDate(0, 0, 7*(x-1))
	if date.Month().String() == month {
		return int(date.Day())
	}

	return -1
}
