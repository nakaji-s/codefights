package main

import "regexp"
import "strings"

func programTranslation(solution string, args []string) string {
	argumentVariants := strings.Join(args, "|")
	re := regexp.MustCompile(`([\s\(\[,!+-])(` + argumentVariants + `)([\s\),;.\]+-])`)
	repl := `$1$$$2$3`
	subSolution := re.ReplaceAllString(solution, repl)
	return re.ReplaceAllString(subSolution, repl)
}
