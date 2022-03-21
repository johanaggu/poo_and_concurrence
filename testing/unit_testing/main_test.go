package main

import (
	"testing"

)

func TestSum(t *testing.T)  {	
	table := []struct{
		a int
		b int
		expected int
	}{
		{4, 5,9,},
		{15,23,38},
		{15,5,20},
	}


	for _, v := range table {
		if result := Sum(v.a, v.b); v.expected != result  {
			t.Errorf("Sum was incorrect, got %d expected %d",result,v.expected)
		}
	}
}
func TestFibonacci(t *testing.T)  {	
	table := []struct{
		number int
		expected int
	}{
		{3, 2},
		{5, 5},
		{6,8},
	}


	for _, v := range table {
		if number := Fibonacci(v.number); v.expected != number  {
			t.Errorf("'Fibonacci' was incorrect, got %d expected %d",number,v.expected)
		}
	}
}