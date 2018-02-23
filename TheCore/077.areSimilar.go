package main

func areSimilar(a []int, b []int) bool {
	unmatched := 0
	unmatchedNum := []int{}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			unmatched++
			if unmatched >= 3 {
				return false
			}
			unmatchedNum = append(unmatchedNum, i)
		}
	}
	if unmatched == 0 {
		return true
	}
	if unmatched == 1 {
		return false
	}
	if a[unmatchedNum[0]] != b[unmatchedNum[1]] {
		return false
	}
	if a[unmatchedNum[1]] != b[unmatchedNum[0]] {
		return false
	}
	return true
}
