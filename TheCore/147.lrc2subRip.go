package main

import (
	"fmt"
	"strconv"
)

func lrc2subRip(lrcLyrics []string, songLength string) []string {
	ret := []string{}
	for i, lrcLyric := range lrcLyrics {
		ret = append(ret, "")
		ret = append(ret, fmt.Sprint(i+1))

		minutes, _ := strconv.ParseInt(lrcLyric[1:3], 10, 32)
		startTime := fmt.Sprintf(`%02d:%02d%s,%s0`, minutes/60, minutes%60, lrcLyric[3:6], lrcLyric[7:9])
		endTime := ""
		if len(lrcLyrics) <= i+1 {
			endTime = songLength + ",000"
		} else {
			minutes, _ := strconv.ParseInt(lrcLyrics[i+1][1:3], 10, 32)
			endTime = fmt.Sprintf(`%02d:%02d%s,%s0`, minutes/60, minutes%60, lrcLyrics[i+1][3:6], lrcLyrics[i+1][7:9])
		}
		time := startTime + " --> " + endTime
		ret = append(ret, time)

		if len(lrcLyric) > 10 {
			ret = append(ret, lrcLyric[11:])
		} else {
			ret = append(ret, "")
		}
	}

	return ret[1:]
}
