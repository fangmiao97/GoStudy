package leetcode

// 376. Wiggle Subsequence
// just to find the rotation of the sequence
// kind of greedy algorithm
func wiggleMaxLength(nums []int) int {
	len_nums := len(nums)
	if len_nums < 2 {
		return len_nums
	}

	res := 0

	// dir == 1 -> up
	// dir == 0 -> down
	var dir int

	for i := 1; i < len_nums; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > nums[i-1] {
			if dir != 1 {
				res++
				dir = 1
			}
			dir = 1
		} else {
			if dir != 2 {
				res++
				dir = 2
			}
		}
	}

	return res + 1
}
