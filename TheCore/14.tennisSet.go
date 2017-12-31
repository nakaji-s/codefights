package main

func tennisSet(score1 int, score2 int) bool {
	if score1 == 7 && score2 < 7 && score2 >= 5 {
		return true
	}
	if score2 == 7 && score1 < 7 && score1 >= 5 {
		return true
	}
	if score1 == 6 && score2 < 5 {
		return true
	}
	if score2 == 6 && score1 < 5 {
		return true
	}
	return false
}
