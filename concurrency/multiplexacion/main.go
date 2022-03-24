package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := time.Second * 4
	d2 := time.Second * 2

	go doSomething(d1, c1, 1)
	go doSomething(d2, c2, 2)

	// DE ESTA FORMA LEEMOS EN ORDEN PORQUE SE BLOQUEA EL PROGRAMA EN LA 19 Y ESPERA A QUE EL CHANNEL 1 RESUELVA//
	// fmt.Println(<-c1 )
	// fmt.Println(<-c2 )

	for i := 0; i < 2; i++ {
		select {
		case channelMessage1:= <-c1:
			fmt.Println(channelMessage1)
		case channelMessage2:= <-c2:
			fmt.Println(channelMessage2)
		}
	}

}

func doSomething(t time.Duration, c chan<- int, param int) {
	time.Sleep(t)
	c <- param
}
