/*
 * @Description: 【问题1】给定一个非空整数数组，除了某个元素出现2次以外，其余每个元素均出现1次。找出那个出现了2次的元素。（请使用map 解决这个问题）
 * @Author: Fang Miao
 * @Created Date: 2021-06-01
 * @LastEditTime: 2021-06-01 17:16:25
 */

package homework

import (
	"fmt"
)

func DuplicateItemCheck(items []string) {
	set := make(map[string]bool) //just like a set
	for _, item := range items {
		if ok := set[item]; ok {
			fmt.Printf("Duplicated Item is " + item + "\n")
		} else {
			set[item] = true
		}
	}
}
