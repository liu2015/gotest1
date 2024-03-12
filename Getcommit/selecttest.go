package getcommit

import "fmt"

func GetSelect() {
	fmt.Println("你好，这是selec")

	test := 10
	for v := range test {
		fmt.Println(v)
	}
}
