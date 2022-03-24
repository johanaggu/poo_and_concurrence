package main

import (
	"fmt"
)

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id  %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Endind with id %d, job%d and fib %d finished\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	tasks := []int{2,34,34,34,43,34,34,43,32} //Value
	nWorkers := 50                 //space per channel [][][z]

	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)  //Aqui se crean 3 workers
	}

	for _, value := range tasks {
		fmt.Println("First")
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}
