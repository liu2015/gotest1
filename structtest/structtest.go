package structtest

type Structtest struct {
	id   int
	name string
	age  string
}

func (str *Structtest) Setstrut() {
	str.id = 1
	str.age = "男"
	str.name = "mingzi"

}

func (str *Structtest) Getstrut() (nametest string) {
	nametest1 := str.name

	return nametest1 + "姓名  性别" + str.age
}
