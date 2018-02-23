package main

func crosswordFormation(words []string) int {
	sum := 0

	// pick 2 from 1st word
	for num, word := range words {
		for i := 0; i < len(word); i++ {
			for j := i + 2; j < len(word); j++ {
				sum += pickSecondWord(i, j, words, num)
			}
		}
	}

	return sum
}

func pickSecondWord(i int, j int, words []string, firstWordNum int) int {
	sum := 0
	wordsWithoutFirst := append([]string{}, words[:firstWordNum]...)
	wordsWithoutFirst = append(wordsWithoutFirst, words[firstWordNum+1:]...)

	for num, word := range wordsWithoutFirst {
		for k := 0; k < len(word)-(j-i); k++ {
			remaining := append([]string{}, wordsWithoutFirst[:num]...)
			remaining = append(remaining, wordsWithoutFirst[num+1:]...)

			tmp := pickThirdAndForthWord(words[firstWordNum][i], word[k], words[firstWordNum][j], word[k+j-i], remaining)
			// if tmp > 0 {
			//     //fmt.Println(words[firstWordNum], word, i, j, k)
			//     for x, rune := range words[firstWordNum] {
			//         if x == i || x == j {
			//                 fmt.Print("[" + string(rune) + "]")
			//         } else {
			//             fmt.Print(string(rune))
			//         }
			//     }
			//     fmt.Print(", ")
			//     for x, rune := range word {
			//         if x == k || x == k + (j - i) {
			//             fmt.Print("[" + string(rune) + "]")
			//         } else {
			//             fmt.Print(string(rune))
			//         }
			//     }
			//     fmt.Println("\n")
			// }
			sum += tmp
		}
	}

	return sum
}

func pickThirdAndForthWord(firstWordFirst, secondWordFirst, firstWordLast, secondWordLast byte, words []string) int {
	sum := 0
	for num, word := range words {
		for i := 0; i < len(word); i++ {
			if word[i] == firstWordFirst {
				for j := i + 2; j < len(word); j++ {
					if word[j] == secondWordFirst {
						remaining := append([]string{}, words[:num]...)
						remaining = append(remaining, words[num+1:]...)

						for k := 0; k < len(remaining[0])-(j-i); k++ {
							if remaining[0][k] == firstWordLast && remaining[0][k+(j-i)] == secondWordLast {
								// for x, rune := range word {
								//     if x == i || x == j {
								//         fmt.Print("[" + string(rune) + "]")
								//     } else {
								//         fmt.Print(string(rune))
								//     }
								// }
								// fmt.Print(", ")
								// for x, rune := range remaining[0] {
								//     if x == k || x == k + (j - i) {
								//         fmt.Print("[" + string(rune) + "]")
								//     } else {
								//         fmt.Print(string(rune))
								//     }
								// }
								// fmt.Println()

								sum++
							}
						}
					}
				}
			}
		}
	}

	return sum
}
