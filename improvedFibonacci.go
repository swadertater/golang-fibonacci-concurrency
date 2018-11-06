package main

import (
	"fmt"
	"runtime"
	"time"
)

type Fibonacci struct {
	num    float64
	answer float64
}

type goRoutineManager struct {
	goRoutineCnt chan bool
}

func (g *goRoutineManager) Run(f func()) {
	select {
	case g.goRoutineCnt <- true:
		go func() {
			f()
			<-g.goRoutineCnt
		}()
	default:
		f()
	}
}

func NewGoRoutineManager(goRoutineLimit int) *goRoutineManager {
	return &goRoutineManager{
		goRoutineCnt: make(chan bool, goRoutineLimit),
	}
}

func newFibonacci(n float64, gm *goRoutineManager) *Fibonacci {

	f := new(Fibonacci)
	f.num = n
	c1 := make(chan float64, 1)
	c2 := make(chan float64, 1)

	if f.num <= 1 {
		f.answer = n
	} else {
		gm.Run(func() {
			fib1 := newFibonacci(n-1, gm)
			c2 <- fib1.answer
		})

		gm.Run(func() {
			fib2 := newFibonacci(n-2, gm)
			c1 <- fib2.answer
		})

		f.answer = <-c2 + <-c1
	}
	close(c1)
	close(c2)

	return f
}

func main() {

	numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	for _, value := range numbers {
		start := time.Now()
		fmt.Println("getting the ", value, " fibonacci number")
		numCpu := runtime.NumCPU()
		gm := NewGoRoutineManager(numCpu)

		f := newFibonacci(value, gm)
		fmt.Println(f.answer)

		end := time.Now()
		totalTime := end.Sub(start)
		fmt.Println("Fibonacci number: ", value, " took ", totalTime, "\n")
	}
	for _, value := range numbers {
		start := time.Now()
		fmt.Println("getting the ", value, " fibonacci number")

		fmt.Println(fib(int(value)))

		end := time.Now()
		totalTime := end.Sub(start)
		fmt.Println("Fibonacci number: ", value, " took ", totalTime, "\n")
	}

}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
