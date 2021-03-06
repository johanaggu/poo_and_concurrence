package main

import (
	"fmt"
	"sync"
	"time"
)



func main() {
	var wg sync.WaitGroup
	c := make(chan interface{}, 5)

	for i := 0; i < 10; i++ {
		c <- i
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan interface{}) {
	defer wg.Done()
	//Dummy process - simulate long process
	fmt.Printf("Started...%v \n", i)
	time.Sleep(time.Second * 4)
	fmt.Printf("Finished...%v \n", <- c)
	
}
