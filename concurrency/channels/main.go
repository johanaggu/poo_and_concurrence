package main

import (
	"fmt"
	"time"
)

func main() {
	// a := make(chan int) //Creating Channel unbuffered

	c := make(chan int, 3) //Creating Channel with Buffer
	c <- 1                 //send number "1" through channel "c"
	c <- 5
	c <- 9

	fmt.Println(<-c) //Saca un valor del channel
	fmt.Println(<-c)

	goFunc(c)
}

func goFunc(c chan int) {
	time.Sleep(time.Second * 2)
	fmt.Println(<-c)
}
