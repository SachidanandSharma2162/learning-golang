package main
import "fmt"

func main()  {
	// var array [5]int
	// array[0]=0
	// array[1]=1
	// array[2]=2
	// array[3]=3
	// array[4]=4

	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("%d ",array[i])
	// }

	// for _, v := range array {
	// 	fmt.Printf("%d ",v)
	// }
	// fmt.Printf("\n")

	// accessing element:

	// fmt.Println(array[3])
	// var ptr *int=&array[0]

	// for i := 0; i < len(array); i++ {
	// 	fmt.Printf("%d ",*(ptr))
	// }
	arr:=[...]int{1,2,3,4,5}
	fmt.Printf("%d\n",len(arr))

	userName:=[...]string{"aman","anuj","shital"}
	newUserName:=[3]string{"aman","anuj","shital"}
	// both the decleartions are valid
	fmt.Printf("\n this is of type%T \n",userName)
	for _, v := range userName {
		fmt.Println(v);
	}
	for _, v := range newUserName {
		fmt.Println(v);
	}
	// to print whole array
	fmt.Println(userName)
	fmt.Println(len(userName))

	// we can also perform the array slicing
	fmt.Println(userName[1:3])
}
