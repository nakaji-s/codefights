package main

func willYou(young bool, beautiful bool, loved bool) bool {
	if young && beautiful && !loved {
		return true
	}
	if loved && (!young || !beautiful) {
		return true
	}
	return false
}
