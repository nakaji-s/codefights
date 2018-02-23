package main

func beautifulText(inputString string, l int, r int) bool {
	tmp := inputString + " "
OUTER:
	for i := l + 1; i <= r+1; i++ {
		if len(tmp)%i != 0 {
			continue
		}

		for j := i; j < len(tmp)-1; j += i {
			if tmp[j-1] != ' ' {
				continue OUTER
			}
		}
		return true
	}

	return false
}
