package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	for i := 0; i < 3; i++ {
		go doSomething(i, &wg)
	}

	wg.Wait() //block and waiting to ending go rutines
}
func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	//Dummy process - simulate long process	
	fmt.Printf("Started...%v \n", i)
	time.Sleep(time.Second * 2)
	fmt.Println("Finished ")
}
