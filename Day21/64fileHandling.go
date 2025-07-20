package main

import (
	"fmt"
	"os"
)
func createFile(filename string){
	file,err:=os.Create(filename)

	if err!=nil{
		fmt.Println("failed to ceeate file",err)
	}else{
		fmt.Println("file creeated successfully!",file)
		defer file.Close()
	}
}
func writeFile(filename string,content string){
	file,err:=os.OpenFile(filename,os.O_APPEND|os.O_WRONLY,0644)

	if err!=nil{
		fmt.Println("Failed to open file",err)
	}else{
		fmt.Println("File opened successfully")
		_,err:=file.WriteString(content)
		if err!=nil{
			fmt.Println("failed to write into the file",err)
		}else{
			fmt.Println("Content written successfully!")
		}
		defer file.Close()
	}
}

func readFile(filename string) string{
	data,err:= os.ReadFile(filename)
	if err!=nil{
		msg:="Error reading file"
		return msg
	}else{
		content:=string(data)
		return content
	}
}
func removeFile(filename string){
	err:=os.Remove(filename)
	if err!=nil{
		fmt.Println("File not deleted",err)
	}else{
		fmt.Println("file deleted successfully")
	}
}
func main() {
	// createFile("./temp1.txt")

	// content:="This is my report on Golang"
	// writeFile("./temp.txt",content)

	// data:=readFile("./temp.txt")
	// fmt.Println(data)

	// removeFile("./temp.txt") 

	// file,_:=os.Open("./temp.txt")
	// fileInfo,_:=file.Stat()

	// fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.Size())
	// fmt.Println(fileInfo.IsDir())
	// fmt.Println(fileInfo.Mode())
	// fmt.Println(fileInfo.ModTime())


}
