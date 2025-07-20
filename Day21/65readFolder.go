package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")

	if err!=nil{
		panic(err)
	}else{
		file,err:=dir.ReadDir(6)
		if err!=nil{
			panic(err)
		}else{
			for _,f := range file{
				fmt.Println(f)
			}
		}
	}
}