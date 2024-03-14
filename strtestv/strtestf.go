package strtestv

import "fmt"

type Structtagvt struct {
	Id   int    `json:"ffo"`
	Age  string `json:"age"`
	Name string `json:"name"`
}

func (str *Structtagvt) Setstrut() {
	str.Id = 1
	str.Age = "女"
	str.Name = "neorng"
	fmt.Println(str.Id, "id", "性别", str.Age, "neriong", str.Name)

}
