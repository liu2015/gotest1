package main

import (
	"fmt"
	getcommit "gotest/Getcommit"
	"gotest/mainfmt"
	"gotest/oraclenumber"
	"time"
)

func main() {
	// fmt.print("你好")
	fmt.Println("你好，这是一个程序的开始")
	getcommit.Getcommit()
	getcommit.Dev()
	getcommit.Getmerge()
	getcommit.GetSelect()

	fmt.Println("递交分支到远程")
	mainfmt.Nametest()
	oraclenumber.Oracletest()

	time.Sleep(5 * time.Second)

}
