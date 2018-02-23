package main

func reflectString(inputString string) string {
	ret := []rune{}
	for _, rune := range inputString {
		ret = append(ret, 'a'+'z'-rune)
	}

	return string(ret)
}
