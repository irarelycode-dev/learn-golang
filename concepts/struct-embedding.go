package main

import "fmt"

func print(args ...interface{}) {
	fmt.Println(args...)
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num %v", b.num)
}

type container struct {
	base
	str string
}

func structEmbedding() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some random string",
	}
	print(co.base.describe())
	print("direct access to num:", co.num)
	print("full path access to num:", co.base.num)
	print("base methods can be accessed directly by co:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	print(d.describe())

}
