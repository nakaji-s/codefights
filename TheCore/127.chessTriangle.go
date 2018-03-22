package main

func chessTriangle(n int, m int) int {
	return 8 * (m0(m-1)*m0(n-2) +
		m0(m-2)*m0(n-1) +
		m0(m-1)*m0(n-3) +
		m0(m-3)*m0(n-1) +
		2*m0(m-2)*m0(n-2) +
		m0(m-2)*m0(n-3) +
		m0(m-3)*m0(n-2))
}

func m0(n int) int {
	if n < 0 {
		return 0
	}
	return n
}
