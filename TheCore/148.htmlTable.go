package main

import "regexp"

func htmlTable(table string, row int, column int) string {
	reTr := regexp.MustCompile(`<tr>(.*?)</tr>`)
	reTd := regexp.MustCompile(`<td>(.*?)</td>`)
	matches := reTr.FindAllStringSubmatch(table, -1)

	m := make([][]string, len(matches))
	for i := 0; i < len(matches); i++ {
		tdMatches := reTd.FindAllStringSubmatch(matches[i][1], -1)
		m[i] = make([]string, len(tdMatches))

		for j := 0; j < len(tdMatches); j++ {
			m[i][j] = tdMatches[j][1]
		}
	}

	if row < len(m) && column < len(m[row]) {
		return m[row][column]
	}
	return "No such cell"
}
