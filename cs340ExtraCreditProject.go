package main

import (
	"fmt"
	"time"
	"runtime"
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
		go func() {
			fib1 := newFibonacci(n - 1)
			c2 <- fib1.answer
		}()
		go func() {
			fib2 := newFibonacci(n - 2)
			c1 <- fib2.answer	
		}()
		
		f.answer = <-c2 + <-c1
	}
	close(c1)
	close(c2)

	return f
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	numbers := []float64{30, 35, 36, 37, 38, 39, 40}
	numProcessors := runtime.NumCPU
	for w := 0; w  numProcessors; w++{
		worker(w)
	}

	
	
	for _, value := range numbers{
		start := time.Now()
		fmt.Println("getting the ", value, " fibonacci number")
		f := newFibonacci(value)
		fmt.Println(f.answer)
		end := time.Now()
		totalTime := end.Sub(start)
		fmt.Println("Fibonacci number: ", value, " took ", totalTime, "\n")
	}
	
}

func fib(n float64) float64 {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
