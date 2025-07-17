package main
import "fmt"

func main(){

	directions:=map[string]int{
		"left":1,
		"right":2,
		"up":3,
		"down":4,
	}

	fmt.Printf("%T\n",directions)
	fmt.Println(directions)

	// accessing values
	first:=directions["left"]
	fmt.Println(first)

	// check if key exist
	 direction,ok:=directions["up"]

	if ok{
		fmt.Println("Found:",direction)
	}else{
		fmt.Println("Not found")
	}

	// deletion in the map
	delete(directions,"down")
	fmt.Println(directions)

	// looping in map
	for key, value := range directions {
		fmt.Printf("%s : %d\n",key,value)
	}

	// clear values
	for k := range directions {
		delete(directions, k)
	}
	fmt.Println(directions)

	for i,c:= range "golang"{
		fmt.Println(i,string(c))
	}
}