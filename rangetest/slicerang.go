package rangetest

import "fmt"

func Slicerang() {
	Slice := []int{1, 23, 4, 231}
	Slice1 := make([]int, 10)
	for _, v := range Slice {
		fmt.Println(v)
	}
	sl := 10
	for v := range sl {
		// 第九次

		fmt.Println("执行第", v)
		Slice1[v] = v
		fmt.Println(Slice1, "默认值 ", Slice1[v])
	}

}
