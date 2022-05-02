package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] {
			res = append(res, prev)
			prev = cur
		} else {
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	res := ""
	target := ""
	for i := 1; i <= len(strs[0]); i++ {
		if i == len(strs[0]) {
			target = strs[0]
		} else {
			target = strs[0][:i]
		}

		for index, v := range strs {
			if len(v) < i {
				break
			}
			current := v[:i]
			if current == target {
				if index == len(strs)-1 {
					res = target
					if i == len(strs[0]) {
						break
					} else {
						target = strs[0][:i+1]
						break
					}
				}
			} else {
				return res
			}
		}
	}
	return res
}
