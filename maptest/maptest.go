package maptest

import "fmt"

func Maptest() {
	m := map[string]int{
		"app":  3,
		"test": 4,
		"ban":  4,
	}

	m["dtd"] = 34

	for k, v := range m {
		fmt.Println(k, v)
	}

	ted := 100000
	for v := range ted {
		fmt.Println("测试遍历：", v)
	}

}
