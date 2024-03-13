package chantest

import (
	"fmt"
	"time"
)

func Chanselect1() {

	chantest1 := make(chan int, 20)
	timeout := time.After(1 * time.Second)

	go func() {
		for v := range 20 {
			// chan管道写入
			fmt.Println("写入chan通道", v)
			chantest1 <- v
		}

	}()

	go func() {
		for {
			select {
			case num1 := <-chantest1:
				fmt.Println("读取出来1", num1)
			// case num2 := <-chantest1:
			// 	fmt.Println("读取出来2", num2)
			case <-timeout:
				fmt.Println("缓存2秒读取")
				fmt.Println("准备退出循环")
				return

			}
		}

	}()

}
