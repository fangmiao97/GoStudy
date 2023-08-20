/*
 * @Description:
 * @Author: Fang Miao
 * @Created Date: Do not edit
 * @LastEditTime: 2022-12-22 06:42:42
 */
package main

import "fmt"

func main() {
	chanel := make(chan struct{})
	receive(chanel)

}

func receive(ch chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Printf("\"get something\": %v\n", "get something")
			v := <-ch
			fmt.Printf("v: %v\n", v)
		}
	}
}
