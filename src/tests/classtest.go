package tests

import "fmt"

type BaseClass struct {
	baseProperty string
}

func NewBaseClass() BaseClass{
	return BaseClass{
		baseProperty: "base property",
	}
}

type SubClass struct {
	BaseClass

	subProperty string
}

func NewSubClass() SubClass {
	s := SubClass{
		subProperty: "sub property",
	}
	if s.baseProperty == "" {
		s.baseProperty = "base property"
	}
	return s
}

func ClassTest() {
	b := NewBaseClass()

	s := NewSubClass()

	fmt.Println(&b)
	fmt.Println(b.baseProperty)

	fmt.Println(&s)
	fmt.Println(s.baseProperty)
	fmt.Println(s.subProperty)
}

