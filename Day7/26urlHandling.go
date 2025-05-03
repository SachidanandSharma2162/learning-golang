package main

import (
	"fmt"
	"net/url"
)

const myUrl string="https://localhost:3000/users?name=sachidanand&userid=ghfy363363h"
func main()  {
	fmt.Println(myUrl)
	
	// parsing url
	result,_:=url.Parse(myUrl) 
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	qparams:=result.Query() // returns a result
	for key, value := range qparams {
		fmt.Println(key,":",value)
	}

	fmt.Println(qparams["name"])
	fmt.Println(qparams["userid"])

	// build your custom url

	myCustomUrl:=&url.URL{
		Scheme: "https",
		Host: "localhost:3000",
		Path: "/users",
		RawQuery: "user=sachi",
	}

	finalUrl:=myCustomUrl.String()
	fmt.Println(finalUrl)
}