package main

import (
	"fmt"
	"gotest/structtest"
	"time"
)

func main() {
	// fmt.print("你好")
	fmt.Println("你好，这是一个程序的开始")
	// getcommit.Getcommit()
	// getcommit.Dev()
	// getcommit.Getmerge()
	// getcommit.GetSelect()

	// fmt.Println("递交分支到远程")
	// mainfmt.Nametest()
	// oraclenumber.Oracletest()
	// chantest.Chantest()
	// chantest.Canselect()
	// fmt.Println("递交")
	// fmt.Println("递交")
	// chantest.Chanselect1()
	// rangetest.Slicerang()
	// rangetest.Rangearray()
	ms := new(structtest.Structtest)
	ms.Setstrut()
	testfd := ms.Getstrut()

	fmt.Println(testfd)
	var p1 structtest.Structtest
	p1.Setstrut()
	p1.Getstrut()
	time.Sleep(10 * time.Second)

}
