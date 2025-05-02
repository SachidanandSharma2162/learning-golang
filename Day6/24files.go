package main

import (
	"fmt"
	"os"
)
func createFile(fileName string) {
    // Create the file
    file, err := os.Create(fileName)
    if err != nil {
        fmt.Println("❌ Error creating file:\n", err)
        return
    }
    defer file.Close()

    fmt.Println("✅ File created successfully:", fileName)
}
func writeIntoFile(fileName string,content string)  {
	file,err:=os.OpenFile(fileName,os.O_APPEND|os.O_WRONLY,0644)

	if err!=nil{
		fmt.Println("Error opening file:", err)
        return
	} else {
		_,err:=file.WriteString(content)
		if err!=nil{
			fmt.Println("Error appending:", err)
        return
		}
	}
	fmt.Println("✏️ Text appended successfully.")
	defer file.Close()
}
func readFile(fileName string) string {
	data,err:=os.ReadFile(fileName)
	if err!=nil{
		message:="Error reading file"
        return message
	} else {
		content:=data
		return string(content)
	}
}

func removeFile(fileName string)  {
	err := os.Remove(fileName)
    if err != nil {
        fmt.Println("Error deleting:", err)
    } else {
        fmt.Println("🗑️ File deleted successfully.")
    }
}
func main()  {
	// createFile("./intro.txt")
	
	// content:="This is my first file: Sachidanand Sharma!!"
	// writeIntoFile("./intro.txt",content)

	// fileContent:=readFile("./intro.txt")
	// fmt.Println(fileContent)

	// removeFile("./intro.txt")
}