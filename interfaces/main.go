package main

import "fmt"

type Greetings interface {
	SayHello() string
}

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func (p *Person) SayHello() string {
	return fmt.Sprintf("Hello! Im %s with age %d", p.name, p.age)
}

func SayHello(g Greetings) string {
	return g.SayHello()
}

func main() {

}
