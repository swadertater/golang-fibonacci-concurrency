package main

import (
	"fmt"
	"sync"
)

type Fibonacci struct {
	num    float64
	answer float64
}

func newFibonacci(n float64) *Fibonacci {

	f := new(Fibonacci)
	f.num = n
	c1 := make(chan float64)
	c2 := make(chan float64)

	if f.num <= 1 {
		f.answer = n
	} else {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			fib1 := newFibonacci(n - 1)
			fmt.Println(fib1.answer)
			c1 <- fib1.answer
			wg.Done()
		}()
		go func() {
			fib2 := newFibonacci(n - 2)
			fmt.Println(fib2.answer)
			c2 <- fib2.answer
			wg.Done()
		}()
		wg.Wait()
		f.answer = <-c2 + <-c1
	}
	close(c1)
	close(c2)

	return f
}

func main() {
	var n float64 = 4
	f := newFibonacci(n)
	fmt.Println(f.answer)

}

func fib(n float64) float64 {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
