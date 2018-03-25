package main

import "time"

func validTime(t string) bool {
	layout := "15:04"
	_, err := time.Parse(layout, t)

	return err == nil
}
