package getcommit

import "fmt"

func GetSelect() {
	fmt.Println("你好，这是selec")

	testd := []int{231, 132, 13232, 12}

	go func() {
		for i, v := range testd {
			fmt.Println("测试", v, "序号", i)
		}
	}()

	test := 10
	for v := range test {
		fmt.Println(v)
	}

}
