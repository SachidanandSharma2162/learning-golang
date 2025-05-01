package main
import "fmt"

func updateValue(x *int) {
	*x = *x + 10
}
func main()  {
	// fmt.Println("Now we will be learning about pointers.")
	// a:= 10
	// p:=&a
	// fmt.Printf("value is:%d\naddress is: %p\n",a,p)
	// fmt.Printf("value is: %d",*(p))

	// a := 42
    // p := &a

    // fmt.Println("Value of a:", a)
    // fmt.Println("Address of a:", p)
    // fmt.Println("Value from pointer:", *p)

    // *p = 100 // changing value through pointer
    // fmt.Println("New value of a:", a)
	
	
	// pass by reference
	// num := 5
	// updateValue(&num)
	// fmt.Println("Updated:", num) // 15
	
<<<<<<< HEAD
	// a := 25
    // p := &a

    // fmt.Printf("Value of a: %d\n", a)
    // fmt.Printf("Address of a: %p\n", &a)
    // fmt.Printf("Pointer p: %p\n", p)
    // fmt.Printf("Value at pointer p: %d\n", *p)

    // // Change value using pointer
    // *p = 50
    // fmt.Printf("New value of a after pointer update: %d\n", a)

	var number int =10
	var ptr *int= &number

	fmt.Printf("value of number : %d \n",number)
	fmt.Printf("adress of number : %p \n",ptr)
	fmt.Println("value of ptr: \n",ptr)
	fmt.Printf("value at ptr: %d\n",*(ptr))

=======
	a := 25
    p := &a

    fmt.Printf("Value of a: %d\n", a)
    fmt.Printf("Address of a: %p\n", &a)
    fmt.Printf("Pointer p: %p\n", p)
    fmt.Printf("Value at pointer p: %d\n", *p)

    // Change value using pointer
    *p = 50
    fmt.Printf("New value of a after pointer update: %d\n", a)
>>>>>>> 174878aa41d41b89ef04d060a7c2525d526e5b8a
}