package leetcode

func pivotIndex(nums []int) int {
	pivotIndex := -1
	if len(nums) == 0 {
		return pivotIndex
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	leftSum := 0
	for i, v := range nums {
		if leftSum == (sum-v)/2 && (sum-v)%2 == 0 {
			pivotIndex = i
			break
		} else {
			leftSum += v
		}
	}
	return pivotIndex

}
