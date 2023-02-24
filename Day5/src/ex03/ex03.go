package ex03

import (
	"day5/present"
)

func GrabPresents(presents []present.Present, capacity int) (res []present.Present) {
	if capacity < 0 {
		return
	}
	n := len(presents)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= capacity; j++ {
			if j < presents[i-1].Size {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-presents[i-1].Size]+presents[i-1].Value)
			}
		}
	}
	i, j := n, capacity
	for i > 0 && j > 0 {
		if dp[i][j] == dp[i-1][j] {
			i--
		} else {
			res = append(res, presents[i-1])
			j -= presents[i-1].Size
			i--
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
