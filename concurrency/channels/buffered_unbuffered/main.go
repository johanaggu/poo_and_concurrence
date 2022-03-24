package main

import "fmt"

func insert(c chan int) {
	c <- 5
}
func main() {
	c := make(chan int, 5)

	go func(c chan int){
		for v := range c {
			fmt.Println(v)
		}
	}(c) 

	fmt.Println(<-c)
}
