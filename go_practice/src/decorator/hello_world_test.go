package main

import "testing"

func TestHello(t *testing.T) {
	decorator(Hello)("hello, world 1.")
	hello := decorator(Hello)
	hello("hello, world 2.")
}
