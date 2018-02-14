package main

func arrayPreviousLess(items []int) []int {
	ret := []int{-1}
	for i := 1; i < len(items); i++ {
		j := 0
		for j = i - 1; j >= 0; j-- {
			if items[i] > items[j] {
				ret = append(ret, items[j])
				break
			}
		}

		if j == -1 {
			ret = append(ret, -1)
		}
	}

	return ret
}
