package main

func switchLights(a []int) []int {
	ret := []int{}
	turnOnCount := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == 1 {
			turnOnCount++
		}

		if turnOnCount%2 == 0 {
			ret = append(ret, a[i])
		} else {
			ret = append(ret, a[i]^1)
		}
	}

	// reverse
	for i := len(ret)/2 - 1; i >= 0; i-- {
		opp := len(ret) - 1 - i
		ret[i], ret[opp] = ret[opp], ret[i]
	}

	return ret
}
