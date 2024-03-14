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

}
