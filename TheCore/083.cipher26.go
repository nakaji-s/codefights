package main

func cipher26(message string) string {
	sum := 0
	ret := []byte{}
	for _, runeValue := range message {
		decryptedRune := (byte((-sum+(int(runeValue)-'a'))%26+26)%26 + 'a')
		ret = append(ret, decryptedRune)
		sum += int(decryptedRune) - 'a'
	}

	return string(ret)
}
