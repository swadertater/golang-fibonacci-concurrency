package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(Rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	go f(1)
	go f(2)
	go f(3)
	var input string
	fmt.Scanln(&input)
}
