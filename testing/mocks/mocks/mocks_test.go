package mocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			1,
			"dni_1",
			func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{1}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{"dni_1", "John", 22}, nil
				}
			},
			FullTimeEmployee{
				Employee{
					1,
				},
				Person{
					"dni_1",
					"John",
					22,
				},
			},
		},
		{
			2,
			"dni_2",
			func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{2}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{"dni_2", "John", 25}, nil
				}
			},
			FullTimeEmployee{
				Employee{
					2,
				},
				Person{
					"dni_2",
					"John",
					25,
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI
	//Running test...
	for _, test := range table {
		test.mockFunc()
		ftEmployee, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Errorf("Error when getting FullTimeEmployee")
		}
		assert.Equal(t, test.expectedEmployee, ftEmployee)
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
