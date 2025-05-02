package main

import (
	"fmt"
)

type Student struct{
	fullName string
	age int
	ontainedMarks int
	totalMarks int
	dob string
}
func (student Student)chnageName()  {
	// pass by value
	student.fullName="Amit"
}
func (student *Student)chnageNameWithPointer()  {
	// pass by reference
	student.fullName="Amit"
}
func (student Student)printResult()  {
	fmt.Printf("The result of %s is: \n",student.fullName)
	fmt.Printf("Age is: %d \n",student.age)
	fmt.Printf("Obtained Marks is: %d \n",student.ontainedMarks)
	fmt.Printf("Toral Marks is: %d \n",student.totalMarks)
	fmt.Printf("Date of Birth is: %s \n",student.dob)
}
func main()  {
	fmt.Println("Methods")

	student:=Student{"Aman",23,467,500,"23-11-2002"}
	student.printResult()
	// student.chnageName()
	student.chnageNameWithPointer()
	student.printResult()

}