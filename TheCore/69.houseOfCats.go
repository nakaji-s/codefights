package main

func houseOfCats(legs int) []int {
	ret := []int{}
	maxCats := legs / 4
	for i := maxCats; i >= 0; i-- {
		ret = append(ret, (legs-i*4)/2)
	}

	return ret
}
