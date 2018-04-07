package main

import "regexp"

func isSubsequence(t string, s string) bool {
	pattern := ""
	for _, ch := range s {
		pattern += regexp.QuoteMeta(string(ch)) + `.*`
	}
	re := regexp.MustCompile(pattern)
	return re.MatchString(t)
}
