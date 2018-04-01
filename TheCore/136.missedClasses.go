package main

import (
	"fmt"
	"time"
)

func missedClasses(year int, daysOfTheWeek []int, holidays []string) int {
	m := map[int]struct{}{}
	for _, num := range daysOfTheWeek {
		m[num] = struct{}{}
		if num == 7 {
			m[0] = struct{}{}
		}
	}

	layout := "01-02-2006"
	count := 0
	for _, holiday := range holidays {
		holiday = holiday + fmt.Sprintf("-%d", year)
		date, _ := time.Parse(layout, holiday)
		if date.Month() < 6 {
			date = date.AddDate(1, 0, 0)
		}
		weekday := int(date.Weekday())

		if _, ok := m[weekday]; ok {
			count++
		}
	}

	return count
}
