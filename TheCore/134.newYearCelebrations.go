package main

import "time"

func newYearCelebrations(takeOffTime string, minutes []int) int {
	layout := "15:04"
	t0, _ := time.Parse(layout, takeOffTime)
	if takeOffTime == "00:00" {
		t0 = t0.AddDate(0, 0, 1)
	}

	t := t0
	ret := 0
	minutes = append([]int{0}, minutes...)

	for i := 1; i < len(minutes); i++ {
		newTime := t.Add(time.Minute * time.Duration(minutes[i]-minutes[i-1]))
		if newTime.Day() == 2 && t.Add(-time.Second).Day() == 1 { // day 2 is celebrate day
			ret++
		}
		t = newTime.Add(-time.Hour)

	}

	// day 2 is celebrate day
	t = t.Add(-time.Second)
	if t.Day() == 1 {
		ret++
	}

	return ret
}
