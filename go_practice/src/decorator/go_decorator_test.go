package main

import (
	"fmt"
	"testing"
)

type MyFoo func(int, int, int) int
type MyBar func(string, string) string

func TestGoDecorator(t *testing.T)  {
	var myFoo MyFoo
	Decorator(&myFoo, foo)
	fmt.Println("result:", myFoo(1, 2, 3))

	fmt.Println("--------分割线--------")

	var myBar MyBar
	Decorator(&myBar, bar)
	fmt.Println("result:", myBar("hello", "world"))

	fmt.Println("--------分割线--------")

	myBar1 := bar
	Decorator(&myBar1, bar)
	fmt.Println("result:", myBar1("hello", "world!"))
}


func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + ", " + b
}