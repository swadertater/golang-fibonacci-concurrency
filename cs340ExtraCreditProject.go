package main

import (
	"fmt"
	"time"
)

func fib(n float64) float64 {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

type Fibonacci struct {
	num    float64
	answer float64
}

func (f Fibonacci) run() {
	if f.num <= 1 {
		f.answer = f.num
	} else {
		var fib1 Fibonacci
		fib1.num = f.num - 1
		var fib2 Fibonacci
		fib2.num = f.num - 2
		go fib(fib2.num)
		f.answer = fib1.answer + fib2.answer
	}
}

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := 12
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
