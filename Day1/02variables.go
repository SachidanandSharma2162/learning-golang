package main
import "fmt"

const loginToken string="vnvnvnvnnv"
// since LoginToken starts with capital L so it is now automatically a public variable, if we use small l as loginToken we can't access it any where outside this file

func main(){
	var age int = 25
    var pi float64 = 3.1415
    var name string = "Arin"
    var isCool bool = true
	fmt.Println(age, pi, name,isCool)

    fmt.Printf("this is of type: %T \n",name)

    // default values and alias
    var accountBalance string
    fmt.Println(accountBalance)
    fmt.Printf("this is of type: %T \n",accountBalance)

    // implicit way to declare variable,if you haven't declared the type of variable then lexer automatically assign based on the value it holds
    var a=10
    fmt.Println(a)

    // no var style
    // this method can only be used inside any method, but you can not use it to declare any global variable
    numberOfVideos := 10
    fmt.Println(numberOfVideos)

    fmt.Printf("type of this is: %T \n",loginToken)
}