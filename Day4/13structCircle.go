package main
import "fmt"

// name of the structure must be in capital latters 

type Circle struct{
	Radius float64
}
func main()  {
	circle:=Circle{Radius:5}
	fmt.Printf("%f\n",3.14*circle.Radius*circle.Radius)
	circles:=[]Circle{
		{Radius: 10},
		{Radius: 3},
		{Radius: 15},
	}
	for _, v := range circles {
		fmt.Printf("%f\n",3.14*v.Radius*v.Radius)
	}
}