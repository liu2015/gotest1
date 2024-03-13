package chantest

import "fmt"

func Chantest() {
	var mych chan int

	mych = make(chan int, 4)
	fmt.Println("长度~~：", len(mych), "容量", cap(mych))

}
