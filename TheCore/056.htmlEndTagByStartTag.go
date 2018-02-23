package main

import (
	"fmt"
	"strings"
)

func htmlEndTagByStartTag(startTag string) string {
	tmp := startTag[1 : len(startTag)-1]
	ret := strings.Split(tmp, ` `)
	return fmt.Sprintf(`</%s>`, ret[0])
}
