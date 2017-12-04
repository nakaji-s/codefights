package main

import "net"

func isIPv4Address(inputString string) bool {
	if net.ParseIP(inputString) != nil {
		return true
	}
	return false
}
