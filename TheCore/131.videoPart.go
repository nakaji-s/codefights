package main

import "time"

func videoPart(part string, total string) []int {
	layout := "15:04:05"
	t1, _ := time.Parse(layout, part)
	t2, _ := time.Parse(layout, total)
	t0 := time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)

	partDuration := t1.Sub(t0).Nanoseconds()
	totalDuration := t2.Sub(t0).Nanoseconds()
	gcd := getGCD(partDuration, totalDuration)

	return []int{int(partDuration / gcd), int(totalDuration / gcd)}
}

func getGCD(a, b int64) int64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}
