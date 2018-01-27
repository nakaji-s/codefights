package main

func isSubstitutionCipher(string1 string, string2 string) bool {
	transformMap := map[byte]byte{}
	transformMapReverse := map[byte]byte{}

	for i := 0; i < len(string1); i++ {
		if _, ok := transformMapReverse[string2[i]]; ok {
			if transformMapReverse[string2[i]] != string1[i] {
				return false
			}
		}

		if num, ok := transformMap[string1[i]]; !ok {
			transformMap[string1[i]] = string2[i]
			transformMapReverse[string2[i]] = string1[i]
		} else {
			if string2[i] != num {
				return false
			}
		}
	}

	return true
}
