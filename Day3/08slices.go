package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Now we will be discussing Slices in Go");

	userName:=[]string{"aman","arun","anurag"}

	for i, v := range userName {
		fmt.Printf("Index %d: %s\n", i, v)
	}
	fmt.Println(len(userName))
	fmt.Println(cap(userName))

	userName=append(userName, "mango")
	fmt.Println(userName)
	
	copy(userName,userName)
	fmt.Println(userName)
	userName = append(userName, "animal","solid")
	fmt.Println(userName)

	newSlice:=make([]string, len(userName))
	copy(newSlice,userName)
	fmt.Println(newSlice)

	// new way to make slice

	s:=make([]int,0,1)
	s = append(s, 10,1,4,8)
	fmt.Println(s)
	fmt.Println(cap(s))
	fmt.Printf("This is of type : %T \n",userName)

	// Slice slicing is also possible
	fmt.Println(userName[1:3])

	// sorting the slice
	sort.Strings(userName)
	fmt.Println(userName)
	// to check is slice is sorted
	fmt.Println(sort.StringsAreSorted(userName))


	fmt.Println(sort.IntsAreSorted(s))
	sort.Ints(s)
	fmt.Println(sort.IntsAreSorted(s))
}