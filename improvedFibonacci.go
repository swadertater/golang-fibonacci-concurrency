package main

import(
	"fmt"
	"runtime"
	"sync"
)

func main(){
	numCpus := runtime.NumCPU
	var wg sync.WaitGroup
	fmt.Println("We have ", numCpus, " available cores.")
	jobs := make(chan int, 20)
	results := make(chan int, 20)

	for i := 0; i < 20; i++ {
		jobs <- i
	}

	

	for w:= 1; w < 3; w++{
		go worker(w, jobs, results)
	}

}
/*
func worker(chan int){

}
*/
func fib(n int) int{
	if n <= 1{
		return n
	}
	return fib(n-1) + fib(n-2)
}

func worker(id int, jobs <-chan int, results chan<-int){
	for j:= range jobs{
		fibonacciNumber := fib(j)
		fmt.Println("Woker: ", id, " has completed the job and returned", fibonacciNumber)
	}
}
