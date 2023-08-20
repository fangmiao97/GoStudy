/*
 * @Description:
 * @Author: Fang Miao
 * @Created Date: Do not edit
 * @LastEditTime: 2022-06-14 15:52:31
 */
package others

import "testing"

func Test_typefun(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Test_typefun"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typefun()
		})
	}
}
