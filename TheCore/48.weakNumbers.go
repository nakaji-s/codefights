package main

func weakNumbers(n int) []int {
	divMap := map[int]int{}
	weaknessMap := map[int]int{}
	maximalWeakness := 0
	maximalWeaknessCount := 0

	for i := 1; i <= n; i++ {
		divMap[i] = calcDiv(i)

		largerCount := 0
		for j := 1; j < n; j++ {
			if divMap[i] < divMap[j] {
				largerCount++
			}
		}
		weaknessMap[i] = largerCount

		if maximalWeakness < largerCount {
			maximalWeakness = largerCount
			maximalWeaknessCount = 1
		} else if maximalWeakness == largerCount {
			maximalWeaknessCount++
		}
	}

	return []int{maximalWeakness, maximalWeaknessCount}
}

func calcDiv(n int) int {
	ret := 0
	for i := n; i > 0; i-- {
		if n%i == 0 {
			ret++
		}
	}

	return ret
}
