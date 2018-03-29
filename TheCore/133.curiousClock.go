package main

import "time"

func curiousClock(someTime string, leavingTime string) string {
	layout := "2006-01-02 15:04"
	t1, _ := time.Parse(layout, someTime)
	t2, _ := time.Parse(layout, leavingTime)

	duration := t1.Sub(t2)
	retTime := t1.Add(duration)

	return retTime.Format(layout)
}
