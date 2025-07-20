// streaming fashion

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	srcFile,err:=os.Open("temp.txt")
	if err!=nil{
		fmt.Println(err)
	}
	defer srcFile.Close()

	destFile,err:=os.Create("temp1.txt")
	if err!=nil{
		fmt.Println(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(srcFile)
	writer:=bufio.NewWriter(destFile)

	for {
		b,err:=reader.ReadByte()
		if err!=nil{
			if err.Error()!="EOF"{
				panic(err)
			}
			break
		}
		e:=writer.WriteByte(b)
		if e!=nil{
            panic(e)
		}
	}
	writer.Flush()
	fmt.Println("File Written SUccessfully")

}