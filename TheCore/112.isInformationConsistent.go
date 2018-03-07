package main

func isInformationConsistent(evidences [][]int) bool {
	for i := 0; i < len(evidences[0]); i++ {
		hasGuilty := false
		hasInnocent := false
		for j := 0; j < len(evidences); j++ {
			if evidences[j][i] > 0 {
				hasGuilty = true
			} else if evidences[j][i] < 0 {
				hasInnocent = true
			}
		}

		if hasGuilty && hasInnocent {
			return false
		}
	}

	return true
}
