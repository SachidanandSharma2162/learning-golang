package main

import "fmt"
func incrementByValue(num int){
	// call by Value
	num=num+1
}
func incrementByReference(num *int){
	// call by Reference
	*num=*num+1
}
func main(){
	var num int=10
	incrementByValue(num)
	fmt.Println(num)
	incrementByReference(&num)
	fmt.Println(num)
}