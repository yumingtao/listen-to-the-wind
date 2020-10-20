package main

import "fmt"

func decorator(f func(string)) func(string) {
	return func(s string) {
		fmt.Println("start...")
		f(s)
		fmt.Println("end...")
	}
}

func Hello(s string) {
	fmt.Println(s)
}
