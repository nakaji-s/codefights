package main

func almostIncreasingSequence(sequence []int) bool {
	count := 0
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i+1]-sequence[i] <= 0 {
			// more than one returns false
			if count > 0 {
				return false
			}

			// calculate again from i-1
			if i > 0 && i+1 < len(sequence)-1 {
				if sequence[i+2]-sequence[i] <= 0 && sequence[i+1]-sequence[i-1] <= 0 {
					return false
				}
			}
			count++
		}
	}
	return true
}
