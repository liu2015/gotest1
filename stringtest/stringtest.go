package stringtest

import "fmt"

func Stringtest() {
	fmt.Println("你好，这是一个软件string")
	defer func() {
		fmt.Println("defer 最后执行")
	}()
	str := "neirong"
	typestr := []byte(str)
	fmt.Println(str)
	fmt.Println(typestr)

	tes := 20
	for v := range tes {
		fmt.Println("tes:", v)
	}

}
