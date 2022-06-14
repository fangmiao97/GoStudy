/*
 * @Description:
 * @Author: Fang Miao
 * @Created Date: Do not edit
 * @LastEditTime: 2022-06-14 15:59:57
 */
package others

import "fmt"

const (
	test_base   = 10000
	test_base_2 = 1000
)

func typefun() {
	var userId int64 = 1200379093
	fmt.Printf("userId: %d\n", userId)
	fmt.Printf("%T\n", (int32(userId)%test_base)/test_base_2)
	fmt.Printf("%T\n", int32(userId%test_base)/test_base_2)
}
