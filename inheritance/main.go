package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) GetMessage() string {
	return fmt.Sprintf("%s with age %d", p.name, p.age)
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func (f *FullTimeEmployee) GetMessage() string {
	return fmt.Sprintf("%s its a full time employee with age %d", f.name, f.age)
}



func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.id = 1
	ftEmployee.name = "Johan"
	ftEmployee.age = 22

	fmt.Println(ftEmployee.GetMessage())
}
