package main

import (
	"fmt"
	"testing"
)

// from fib_test.go
func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(1)
	}
	//fmt.Print("HELLO")
	fmt.Println(b)

}

func BenchmarkNewRedisClient(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewRedisClient()
	}
}
