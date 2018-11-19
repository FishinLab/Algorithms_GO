package main

import (
	"algorithm"
	"fmt"
	"tests"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	language string
}

//1
func (c *Cat) Speak() string {
	return c.language
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func test() {
	cat := Cat{language: "Meow"}
	fmt.Println(cat.Speak())

	animals := []Animal{Dog{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

func main() {
	test()
	tests.ClassTest()

	//algorithm.TwoSum()
	//algorithm.GetYanghuiTrangle()
	//algorithm.FullPermutation()
	//datastruct.TestBuild()
	//algorithm.AddTwoNumber()
	algorithm.RotateImage()
}
