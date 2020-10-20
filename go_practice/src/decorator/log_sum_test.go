package main

import (
	"fmt"
	"testing"
)

func TestLogSum(t *testing.T){
	sum1 := timeSumFunc(Sum1)
	sum2 := timeSumFunc(Sum2)
	fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))
}