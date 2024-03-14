package rangetest

import "fmt"

func Rangearray() {
	array := [5]int{1, 32, 44, 322, 12}
	s1 := array[0:4]
	s2 := array[1:]
	fmt.Println(" xianshi", s1)
	fmt.Println("xianshi ,", s2)
}
