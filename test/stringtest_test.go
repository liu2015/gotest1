package test

import (
	"fmt"
	"gotest/stringtest"
	"gotest/strtestv"
	"testing"
)

func TestString(t *testing.T) {

	stringtest.Stringtest()
}

func TestStrtestf(t *testing.T) {
	df := new(strtestv.Structtagvt)
	df.Setstrut()
	fmt.Println("~~~~~", df.Name)

}
