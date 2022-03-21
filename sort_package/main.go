package main

import (
	"fmt"
	"sort"
)

func main()  {
	var i sort.IntSlice = []int{8,2,6,4,8,2,4,6}
	fmt.Println(i) // before sorted
	sort.Ints(i)
	fmt.Println(i) // After sorted

	
}