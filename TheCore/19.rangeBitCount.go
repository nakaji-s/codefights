package main

func rangeBitCount(a int, b int) int {
	num := int64(0)
	for i := a; i <= b; i++ {
		num += countBits(int64(i))
	}
	return int(num)
}

func countBits(bits int64) int64 {
	bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
	bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
	bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
	bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
	return (bits & 0x0000ffff) + (bits >> 16 & 0x0000ffff)
}
