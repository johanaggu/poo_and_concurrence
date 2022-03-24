package main

import (
	"fmt"
)

func Generator(c chan int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c) //Close channel and block inputs
}

func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value * 2
	}

	close(out)
}

func Print(c chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	double := make(chan int)

	go Generator(generator)
	go Double(generator, double)

	Print(double)
}
