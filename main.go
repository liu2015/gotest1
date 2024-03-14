package main

import (
	"fmt"
	"gotest/maptest"
	"gotest/strtestv"
	"gotest/structtest"
	"reflect"
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
	maptest.Maptest()
	tes := new(strtestv.Structtagvt)
	tes.Setstrut()
	reflect1 := reflect.TypeOf(*tes)
	// var sd strtestv.Structtagvt

	find1 := reflect1.Field(0)
	fmt.Println("这是", find1.Tag.Get("json"), find1.Name, find1.Tag, find1.PkgPath)

	// fmt.Println("获得索引的的信息", find1.Tag.Get("json"))
	v := reflect.ValueOf(*tes)

	for i := 0; i < reflect1.NumField(); i++ {
		f := reflect1.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(f.Name, f.Type, "获得字段内置", val)

	}

	fmt.Println("获得字段值~~", v.Field(1).Interface())

	time.Sleep(10 * time.Second)

}
