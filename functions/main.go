package main

import "fmt"

func suma(ints ...int) int {
	var result int = 0

	for _, v := range ints  {
		result+=v	
	}

	return result
}

func main()  {
	x := 5
	// double x
	y := func () int  {
		return x*2
	}()
	fmt.Println(fmt.Sprintf("Double %v -->", x), y)
	
	//SUMA function

	result := suma(9,5,7,2)
	fmt.Println("suma(9,5,7,2) -->", result)

} 