package chessboard

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a int,b int) int{
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int{
	if a > b {
		return a
	}
	return b
}