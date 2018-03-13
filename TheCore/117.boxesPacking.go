package main

import "sort"

func boxesPacking(length []int, width []int, height []int) bool {
	boxCount := len(length)
	boxes := [][]int{}

	for i := 0; i < boxCount; i++ {
		box := []int{length[i], width[i], height[i]}
		sort.Ints(box)
		boxes = append(boxes, box)
	}

	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i][0] < boxes[j][0]
	})

	for i := 0; i < boxCount-1; i++ {
		if boxes[i][0] >= boxes[i+1][0] ||
			boxes[i][1] >= boxes[i+1][1] ||
			boxes[i][2] >= boxes[i+1][2] {
			return false
		}
	}

	return true
}
