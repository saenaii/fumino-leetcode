package dynamic_programming

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func stringInSlice(arr []string, key string) bool {
	for _, v := range arr {
		if v == key {
			return true
		}
	}
	return false
}
