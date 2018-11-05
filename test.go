package main

import(
	"fmt"
	"time"
)

func main(){

	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		c1 <- "hello"
	}()
	word1 := <-c1
	go func(){
		time.Sleep(time.Millisecond * 5000)
		c2 <- "world"
	}()

	
	word2 := <-c2

	
	fmt.Println(word2)

}