package main
import "fmt"

func main()  {
	fmt.Println("Now we will be learning slice operation.")

	courses:=[]string{"java","python","go","cpp","c","javascript","reactjs"}
	fmt.Println(courses)

	var index int=2 // go will be removed.
	courses=append(courses[:index],courses[index+1:]... )
	fmt.Println(courses)
}
