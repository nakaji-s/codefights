package main

import "regexp"

func eyeRhyme(pairOfLines string) bool {
	re := regexp.MustCompile(`.*(...)\t.*(...)`)
	match := re.FindStringSubmatch(pairOfLines)
	return match[1] == match[2]
}
