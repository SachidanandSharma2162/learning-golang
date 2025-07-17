package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Enter total marks:")
	var total int
	fmt.Scanf("%d",&total)

	switch{
	case total>90:
		fmt.Println("Grade A")
		fallthrough
	case total>75  && total<=90:
		fmt.Println("Grade B")
	case total>60  && total<=75:
		fmt.Println("Grade C")
	case total>45  && total<=60:
		fmt.Println("Grade D")
	case total>30 && total<=45:
		fmt.Println("Grade E")
	default:
		fmt.Println("Fail Grade F")
	}

   whoAmI:= func (i interface{})  {
	switch i.(type){
	case int:
		fmt.Println("Interger")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Boolean")
	default:
		fmt.Println("other")
	}
   }
   whoAmI(3)
}