package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	reader:=bufio.NewReader(os.Stdin)
	number,_:=reader.ReadString('\n')
	fmt.Println(number)
	// type conversion
	// updatedNumber,err:=strconv.ParseFloat(strings.TrimSpace(number),64)
	updatedNumber,err:=strconv.ParseInt(strings.TrimSpace(number),10,64)

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(updatedNumber+1)
	}
}