package chantest

import (
	"fmt"
	"time"
)

func Canselect() {
	canselect := make(chan int, 10)

	nu := 10

	go func() {
		for v := range nu {
			fmt.Println("写入到chan", v)
			// canselect = <-v
			// 写入chan管道
			canselect <- v
			time.Sleep(1 * time.Second)

		}
	}()

	go func() {
		for v := range nu {
			//读取的次数
			unsum := <-canselect

			fmt.Println("读取chan", unsum, "这是第一次读取", v)
			time.Sleep(1 * time.Second)

		}
	}()
}
