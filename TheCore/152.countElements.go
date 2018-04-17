package main

func countElements(inputString string) int {
	quoted := false
	count := 0

	if inputString != "[]" {
		count = 1
	}

	for _, character := range inputString {
		switch byte(character) {
		case ',':
			if quoted == false {
				count++
			}
		case '"':
			if quoted == false {
				quoted = true
			} else {
				quoted = false
			}
		}
	}

	return count
}
