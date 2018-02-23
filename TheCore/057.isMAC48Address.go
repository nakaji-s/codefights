package main

import "net"

func isMAC48Address(inputString string) bool {
	_, err := net.ParseMAC(inputString)
	return err == nil
}
