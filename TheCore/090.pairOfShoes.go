package main

func pairOfShoes(shoes [][]int) bool {
	m := map[int]int{}
	for _, num := range shoes {
		if num[0] == 0 {
			m[num[1]]++
		} else {
			m[num[1]]--
		}
	}

	for _, num := range m {
		if num != 0 {
			return false
		}
	}

	return true
}
