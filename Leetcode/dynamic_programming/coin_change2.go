package leetcode

func change(amount int, coins []int) int {
	// initialize dp array
	dp := make([]int, amount+1)

	dp[0] = 1 // means if target amount - coins price == 0, then there is 1 way to change

	for _, coin := range coins {
		for i := 1; i < amount+1; i++ {
			if i-coin >= 0 { // if current amount got current coin to change
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[amount]
}
