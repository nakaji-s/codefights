package main

func boxBlur(a [][]int) [][]int {
	ret := make([][]int, len(a)-2)
	for i := 0; i < len(a)-2; i++ {
		ret[i] = make([]int, len(a[i])-2)
	}
	for i := 1; i < len(a)-1; i++ {
		for j := 1; j < len(a[i])-1; j++ {
			ret[i-1][j-1] = (a[i-1][j-1] + a[i][j-1] + a[i+1][j-1] +
				a[i-1][j] + a[i][j] + a[i+1][j] +
				a[i-1][j+1] + a[i][j+1] + a[i+1][j+1]) / 9
		}
	}
	return ret
}
