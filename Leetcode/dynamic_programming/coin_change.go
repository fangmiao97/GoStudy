package leetcode

// dynamic programming
func coinChange(coins []int, amount int) int {
	// here ths best result is min amount
	// then the initial value of dp array should be some big value
	// here we use amount+1

	// intialize dp array
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0

	// get answer from 1
	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = getMin(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
