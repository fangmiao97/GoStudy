package leetcode

// 300. Longest Increasing Subsequence
// loop for 2 round
func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	len_nums := len(nums)
	dp := make([]int, len_nums)
	dp[0] = 1

	res := 1

	for i := 1; i < len_nums; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = getMax(dp[i], dp[j]+1)
			}
		}
		res = getMax(res, dp[i])
	}

	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
