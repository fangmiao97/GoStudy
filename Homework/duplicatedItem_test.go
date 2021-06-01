/*
 * @Description: 【问题1】给定一个非空整数数组，除了某个元素出现2次以外，其余每个元素均出现1次。找出那个出现了2次的元素。（请使用map 解决这个问题）
 * @Author: Fang Miao
 * @Created Date: 2021-06-01
 * @LastEditTime: 2021-06-01 17:08:55
 */

package homework

import (
	"testing"
)

func TestDuplicatedItemCheck(t *testing.T) {
	// var items []string = make([]string, 4)
	items := []string{"a", "b", "vc", "a"}
	DuplicateItemCheck(items)
}
