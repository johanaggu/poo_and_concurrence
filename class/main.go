package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id,
		name,
		vacation,
	}
}

func main() {
	// Example 1
	e := Employee{}
	e.SetId(1)
	e.SetName("Johan")

	// Example 2
	e2 := Employee{
		id:       2,
		name:     "Ivan",
		vacation: true,
	}

	// Example 3
	e3 := new(Employee)
	
	// Example 4
	e4 := NewEmployee(21, "Name", false)

	fmt.Println(e.GetName())
	fmt.Println(e2)
	fmt.Println(e3)
	fmt.Println(e4)
}
