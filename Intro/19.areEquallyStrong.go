package main

func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	if yourLeft+yourRight == friendsRight+friendsLeft {
		if yourLeft == friendsLeft || yourLeft == friendsRight {
			return true
		}
	}
	return false
}
